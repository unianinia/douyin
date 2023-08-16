package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/publish/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/publish"
	"douyin/kitex_gen/user"
	"sync"
)

type PublishInfoService struct {
	ctx context.Context
}

func NewPublishInfoService(ctx context.Context) *PublishInfoService {
	return &PublishInfoService{
		ctx: ctx,
	}
}

func (s *PublishInfoService) PublishInfo(req *publish.PublishInfoRequest) (*common.Video, error) {
	var video common.Video

	v, err := db.GetVideoById(req.VideoId)
	if err != nil {
		return &video, err
	}

	video.Id = v.ID
	video.Title = v.Title
	video.PlayUrl = v.PlayURL
	video.CoverUrl = v.CoverURL

	errChan := make(chan error, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		count, e := rpc.CommentCount(s.ctx, video.Id)
		if e != nil {
			errChan <- e
		} else {
			video.CommentCount = count
		}
	}()

	go func() {
		defer wg.Done()
		info, e := rpc.UserInfo(s.ctx, &user.UserInfoRequest{
			CurrentUserId: req.CurrentUserId,
			UserId:        v.AuthorID,
		})
		if e != nil {
			errChan <- e
		} else {
			video.Author = info.User
		}
	}()

	wg.Wait()
	select {
	case <-errChan:
		return &video, err
	default:
	}

	return &video, err
}
