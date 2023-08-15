package service

import (
	"context"
	"douyin/cmd/publish/dal/db"
	"douyin/cmd/publish/dal/oss"
	"douyin/cmd/publish/pkg"
	"douyin/kitex_gen/publish"
	"douyin/pkg/constants"
	"douyin/pkg/ffmpeg"
	"douyin/pkg/minio"
	"fmt"
	"strings"
	"time"
)

type PublishActionService struct {
	ctx context.Context
}

func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{
		ctx: ctx,
	}
}

func (s *PublishActionService) PublishAction(req *publish.PublishActionRequest) (bool, error) {
	var video db.Video
	var f pkg.FilePair
	var err error

	video.AuthorID = req.UserId
	video.Title = req.Title
	video.PublishTime = time.Now()
	f.VideoName = fmt.Sprintf("%d_%d.mp4", req.UserId, video.PublishTime.Unix())
	f.VideoData = req.Data

	f.ImageData, err = ffmpeg.GetSnapShot(req.Data)
	if err != nil {
		return false, err
	}
	f.ImageName = strings.Replace(f.VideoName, "mp4", "jpg", 1)

	oss.UploadFile(f)
	iu, err := minio.GetObjectURL(s.ctx, constants.MinioImageBucketName, f.ImageName)
	vu, err := minio.GetObjectURL(s.ctx, constants.MinioVideoBucketName, f.VideoName)
	video.CoverURL = iu.String()
	video.PlayURL = vu.String()

	_, err = db.CreateVideo(&video)

	return true, err
}
