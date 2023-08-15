package db

import (
	"context"
	"fmt"
	"testing"
)

func TestAddNewMessage(t *testing.T) {
	Init()
	msg := Message{
		ToUserId:   1002,
		FromUserId: 1003,
		Content:    "Hello World",
	}

	ok, err := AddNewMessage(context.Background(), &msg)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(ok)
}

func TestGetLatestMessageByIdPair(t *testing.T) {
	Init()
	msg, err := GetLatestMessageByIdPair(1002, 1003)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", *msg)
}
