package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/publish/dal/db"
	"douyin/kitex_gen/common"
	"douyin/kitex_gen/publish"
	"douyin/kitex_gen/user"
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

	resp, err := rpc.UserInfo(s.ctx, &user.UserInfoRequest{
		CurrentUserId: req.CurrentUserId,
		UserId:        v.AuthorID,
	})
	if err != nil {
		return &video, err
	}

	video.Author = resp.UserInfo

	return &video, err
}
