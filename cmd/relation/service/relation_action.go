package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/relation/dal/mq"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"
	"strconv"
	"strings"
)

const (
	FOLLOW   = 1
	UNFOLLOW = 2
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService 创建用户关系服务
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{
		ctx: ctx,
	}
}

func (s *RelationActionService) RelationAction(req *relation.RelationActionRequest) error {
	if req.ActionType != FOLLOW && req.ActionType != UNFOLLOW {
		return errno.ParamErr
	}

	if req.ToUserId == req.CurrentUserId {
		return errno.ParamErr
	}

	exist, err := rpc.UserExist(s.ctx, req.ToUserId)
	if err != nil {
		return err
	}
	if !exist {
		return errno.UserIsNotExistErr
	}

	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(req.CurrentUserId)))
	sb.WriteString("&")
	sb.WriteString(strconv.Itoa(int(req.ToUserId)))
	sb.WriteString("&")
	sb.WriteString(strconv.Itoa(int(req.ActionType)))

	return mq.AddActor.Publish(s.ctx, sb.String())
}
