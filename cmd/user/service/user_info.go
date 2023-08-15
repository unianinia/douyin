package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/user"
	"sync"
)

type UserInfoService struct {
	ctx context.Context
}

// NewUserInfoService 创建用户登录服务
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

func (s *UserInfoService) UserInfo(req *user.UserInfoRequest) (*common.User, error) {
	userInfo := new(common.User)
	u, err := db.QueryUserInfoById(s.ctx, req.UserId)
	if err != nil {
		return userInfo, err
	}

	userInfo.Id = u.ID
	userInfo.Name = u.UserName

	errChan := make(chan error, 5)
	defer close(errChan)

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		favoriteCount, favoritedCount, e := rpc.FavoriteCount(s.ctx, u.ID)
		if e != nil {
			errChan <- e
		} else {
			userInfo.FavoriteCount = &favoriteCount
			userInfo.TotalFavorited = &favoritedCount
		}
	}()

	go func() {
		defer wg.Done()
		count, e := rpc.PublishCount(s.ctx, u.ID)
		if e != nil {
			errChan <- e
		} else {
			userInfo.WorkCount = &count
		}
	}()

	go func() {
		defer wg.Done()
		if req.CurrentUserId == 0 {
			userInfo.IsFollow = false
		} else {
			isFollow, _, e := rpc.RelationExist(s.ctx, req.CurrentUserId, req.UserId)
			if e != nil {
				errChan <- e
			} else {
				userInfo.IsFollow = isFollow
			}
		}
	}()

	go func() {
		defer wg.Done()
		favoriteCount, favoritedCount, e := rpc.FavoriteCount(s.ctx, req.UserId)
		if e != nil {
			errChan <- e
		} else {
			userInfo.FavoriteCount = &favoriteCount
			userInfo.TotalFavorited = &favoritedCount
		}
	}()

	go func() {
		defer wg.Done()
		followCount, followerCount, e := rpc.RelationCount(s.ctx, req.UserId)
		if e != nil {
			errChan <- e
		} else {
			userInfo.FollowCount = &followCount
			userInfo.FollowerCount = &followerCount
		}
	}()

	wg.Wait()
	select {
	case result := <-errChan:
		return userInfo, result
	default:
	}
	return userInfo, nil
}
