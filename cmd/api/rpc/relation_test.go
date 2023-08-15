package rpc

import (
	"context"
	"douyin/kitex_gen/relation"
	"fmt"
	"testing"
)

func TestRelationAction(t *testing.T) {
	InitRPC()

	resp, err := RelationAction(context.Background(), &relation.RelationActionRequest{
		CurrentUserId: 1001,
		ToUserId:      1002,
		ActionType:    1,
	})

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("%#v\n", resp)
}

func TestRelationCount(t *testing.T) {
	InitRPC()

	c1, c2, err := RelationCount(context.Background(), 1001)

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(c1, c2)
}

func TestRelationExist(t *testing.T) {
	InitRPC()

	e1, e2, err := RelationExist(context.Background(), 1001, 1002)

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(e1, e2)
}

func TestRelationFollowList(t *testing.T) {
	InitRPC()

	resp, err := RelationFollowList(context.Background(), &relation.RelationFollowListRequest{
		UserId:        1002,
		CurrentUserId: 1001,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, u := range resp.UserList {
		fmt.Printf("%#v\n", u)
	}
}

func TestRelationFollowerList(t *testing.T) {
	InitRPC()

	resp, err := RelationFollowerList(context.Background(), &relation.RelationFollowerListRequest{
		UserId:        1001,
		CurrentUserId: 1002,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, u := range resp.UserList {
		fmt.Printf("%#v\n", u)
	}
}

func TestRelationFriendList(t *testing.T) {
	InitRPC()

	resp, err := RelationFriendList(context.Background(), &relation.RelationFriendListRequest{
		UserId:        1001,
		CurrentUserId: 1002,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	for _, u := range resp.UserList {
		fmt.Printf("%#v\n", u)
	}
}
