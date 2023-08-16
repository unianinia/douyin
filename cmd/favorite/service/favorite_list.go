package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

func (s *FavoriteListService) FavoriteList(req *favorite.FavoriteListRequest) ([]*common.Video, error) {
	var videos []*common.Video

	exist, err := rpc.UserExist(s.ctx, req.ToUserId)
	if err != nil {
		return videos, err
	}
	if !exist {
		return videos, errno.UserIsNotExistErr
	}

	// TODO: 完善该逻辑 存在潜在的线程阻塞稳定

	return videos, nil
}
