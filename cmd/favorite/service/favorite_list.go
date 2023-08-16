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

	//videoIds, err := db.GetUserFavoriteIdList(s.ctx, req.ToUserId)
	//if err != nil {
	//	return videos, err
	//}
	//
	//for _, id := range videoIds {
	//	video, err := rpc.PublishInfo(s.ctx, id)
	//	if err != nil {
	//		return videos, err
	//	}
	//	videos = append(videos, video)
	//}
	//
	//num := len(videoIds)
	//videoChan := make(chan common.Video, num)
	//errChan := make(chan error, num)
	//doneChan := make(chan struct{})
	//
	//var wg sync.WaitGroup
	//
	//go func() {
	//	for {
	//		select {
	//		case video := <-videoChan:
	//			videos = append(videos, &video)
	//		case <-doneChan:
	//			return
	//		}
	//	}
	//}()
	//
	//for _, id := range videoIds {
	//	wg.Add(1)
	//	go func(videoId int64) {
	//		defer wg.Done()
	//		video, e := rpc.PublishInfo(s.ctx, videoId)
	//		if e != nil {
	//			errChan <- e
	//		} else {
	//			videoChan <- *video
	//		}
	//	}(id)
	//}
	//
	//wg.Wait()
	//doneChan <- struct{}{}
	//
	//select {
	//case err = <-errChan:
	//	return videos, err
	//default:
	//}

	return videos, nil
}
