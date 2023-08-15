package amqpclt

import (
	"context"
	"strconv"
	"strings"

	"douyin/cmd/relation/dal/db"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

// RelationActionAdd 添加用户消息记录
func (a *Actor) RelationActionAdd(ctx context.Context, relations <-chan amqp.Delivery) {
	for d := range relations {
		// 取出用户ID
		params := strings.Split(string(d.Body), "&")
		userId, err := strconv.Atoi(params[0])
		if err != nil {
			klog.Errorf("transform error：(%v)", err)
		}
		followerId, err := strconv.Atoi(params[1])
		if err != nil {
			klog.Errorf("transform error：(%v)", err)
		}
		action := params[2]

		klog.Infof("relation db option(%v,%v,%v)", userId, followerId, action)
		if action == "1" {
			if exist, _ := db.CheckRelationFollowExist(ctx, int64(followerId), int64(userId)); exist {
				return
			}
			if _, err = db.AddNewRelation(context.Background(), &db.Relation{
				UserId:     int64(userId),
				FollowerId: int64(followerId),
			}); err != nil {
				klog.Errorf("add new relation to db：(%v)", err)
			}
		} else {
			if exist, _ := db.CheckRelationFollowExist(ctx, int64(followerId), int64(userId)); !exist {
				return
			}
			if _, err = db.AddNewRelation(context.Background(), &db.Relation{
				UserId:     int64(userId),
				FollowerId: int64(followerId),
			}); err != nil {
				klog.Errorf("add new relation to db：(%v)", err)
			}
		}

	}
}
