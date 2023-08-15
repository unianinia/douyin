package service

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"
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

func (s *RelationActionService) RelationAction(req *relation.RelationActionRequest) (bool, error) {
	if req.ActionType != FOLLOW && req.ActionType != UNFOLLOW {
		return false, errno.ParamErr
	}

	if req.ToUserId == req.CurrentUserId {
		return false, errno.ParamErr
	}

	exist, err := rpc.UserExist(s.ctx, req.ToUserId)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, errno.UserIsNotExistErr
	}

	r := &db.Relation{
		UserId:     req.ToUserId,
		FollowerId: req.CurrentUserId,
	}

	if req.ActionType == FOLLOW {
		exist, _ = db.CheckRelationFollowExist(s.ctx, r.FollowerId, r.UserId)
		if exist {
			return false, errno.FollowRelationAlreadyExistErr
		}
		return db.AddNewRelation(r)
	} else {
		exist, _ = db.CheckRelationFollowExist(s.ctx, r.UserId, r.FollowerId)
		if !exist {
			return false, errno.FollowRelationNotExistErr
		}
		return db.DeleteRelation(r)
	}
}
