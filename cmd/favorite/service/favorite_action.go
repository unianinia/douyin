package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{
		ctx: ctx,
	}
}

func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionRequest) error {
	exist, _ := rpc.UserExist(s.ctx, req.UserId)
	if !exist {
		return errno.UserIsNotExistErr
	}
	exist, _ = rpc.PublishExist(s.ctx, req.VideoId)
	if !exist {
		return errno.VideoIsNotExistErr
	}

	if req.ActionType == 1 {
		ok, err := db.AddNewFavorite(s.ctx, &db.Favorites{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		})
		if err != nil {
			return err
		}
		if !ok {
			return errno.FavoriteActionErr
		}
	} else {
		ok, err := db.DeleteFavorite(s.ctx, &db.Favorites{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		})
		if err != nil {
			return err
		}
		if !ok {
			return errno.FavoriteActionErr
		}
	}
	return nil
}
