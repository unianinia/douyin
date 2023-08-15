package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/publish/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/publish"
	"douyin/kitex_gen/user"
	"sync"
)

type PublishVideoListService struct {
	ctx context.Context
}

func NewPublishVideoListService(ctx context.Context) *PublishVideoListService {
	return &PublishVideoListService{
		ctx: ctx,
	}
}

func (s *PublishVideoListService) PublishVideoList(req *publish.PublishVideoListRequest) ([]*common.Video, error) {
	var videos []*common.Video

	dbVideos, err := db.GetVideoListByVideoIDList(req.VideoIds)
	if err != nil {
		return videos, err
	}

	videoChan := make(chan common.Video, len(dbVideos))
	errChan := make(chan error, len(dbVideos))
	doneChan := make(chan struct{})

	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case v := <-videoChan:
				videos = append(videos, &v)
			case <-doneChan:
				return
			}
		}
	}()

	for _, v := range dbVideos {
		wg.Add(1)
		go func(video *db.Video) {
			defer wg.Done()
			count, e := rpc.CommentCount(s.ctx, &comment.CommentCountRequest{
				VideoId: video.ID,
			})
			if e != nil {
				errChan <- e
				return
			}

			resp, e := rpc.UserInfo(s.ctx, &user.UserInfoRequest{
				CurrentUserId: video.AuthorID,
				UserId:        video.AuthorID,
			})

			videoChan <- common.Video{
				Id:           video.ID,
				PlayUrl:      video.PlayURL,
				CoverUrl:     video.CoverURL,
				Title:        video.Title,
				CommentCount: count,
				Author:       resp.UserInfo,
			}

		}(v)
	}

	wg.Wait()
	doneChan <- struct{}{}

	select {
	case err = <-errChan:
		return videos, err
	default:
	}

	return videos, nil
}
