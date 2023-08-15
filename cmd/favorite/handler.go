package main

import (
	"context"
	"douyin/cmd/favorite/pkg"
	"douyin/cmd/favorite/service"
	"douyin/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp = new(favorite.FavoriteActionResponse)

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg

	return
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp = new(favorite.FavoriteListResponse)

	videos, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	r := pkg.BuildBaseResp(err)
	resp.StatusCode = r.StatusCode
	resp.StatusMsg = r.StatusMsg
	resp.VideoList = videos

	return
}

// FavoriteCount implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteCount(ctx context.Context, req *favorite.FavoriteCountRequest) (resp *favorite.FavoriteCountResponse, err error) {
	resp = new(favorite.FavoriteCountResponse)

	favoriteCount, favoritedCount, err := service.NewFavoriteCountService(ctx).FavoriteCount(req)
	resp.FavoriteCount = favoriteCount
	resp.FavoritedCount = favoritedCount

	return
}

// FavoriteExist implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteExist(ctx context.Context, req *favorite.FavoriteExistRequest) (resp *favorite.FavoriteExistResponse, err error) {
	resp = new(favorite.FavoriteExistResponse)

	exist, err := service.NewFavoriteExistService(ctx).FavoriteExist(req)
	resp.Exist = exist
	return
}

// FavoriteCountOfVideo implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteCountOfVideo(ctx context.Context, req *favorite.FavoriteCountOfVideoRequest) (resp *favorite.FavoriteCountOfVideoResponse, err error) {
	resp = new(favorite.FavoriteCountOfVideoResponse)

	count, exist, err := service.NewFavoriteCountOfVideoService(ctx).FavoriteCountOfVideo(req)
	resp.FavoritedCount = count
	resp.IsFavorite = exist

	return
}
