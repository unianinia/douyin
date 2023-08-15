package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/comment/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{
		ctx: ctx,
	}
}

func (s *CommentListService) CommentList(req *comment.CommentListRequest) ([]*comment.Comment, error) {
	var comments []*comment.Comment

	exist, _ := rpc.PublishExist(s.ctx, req.VideoId)
	if !exist {
		return comments, errno.VideoIsNotExistErr
	}
	dbComments, err := db.GetCommentListByVideoID(s.ctx, req.VideoId)
	if err != nil {
		return comments, errno.CommentIsNotExistErr
	}

	info, err := rpc.UserInfo(s.ctx, &user.UserInfoRequest{
		CurrentUserId: 0,
		UserId:        req.UserId,
	})
	if err != nil {
		return comments, err
	}

	for _, c := range dbComments {
		comments = append(comments, &comment.Comment{
			Id:         c.ID,
			User:       info.User,
			Content:    &c.CommentText,
			CreateDate: c.CreatedAt.Format("01-02"),
		})
	}
	return comments, err
}
