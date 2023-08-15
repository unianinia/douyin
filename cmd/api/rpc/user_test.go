package rpc

import (
	"context"
	"douyin/kitex_gen/user"
	"fmt"
	"testing"
)

func TestUserRegister(t *testing.T) {
	InitRPC()

	resp, err := UserRegister(context.Background(), &user.UserRegisterRequest{
		Username: "hwx",
		Password: "123456",
	})
	if err != nil {
		return
	}
	fmt.Printf("%#v\n", resp)
}

func TestUserInfo(t *testing.T) {
	InitRPC()

	info, err := UserInfo(context.Background(), &user.UserInfoRequest{
		CurrentUserId: 1001,
		UserId:        1002,
	})
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	fmt.Printf("%#v\n", *info)
}
