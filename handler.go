package main

import (
	"context"
	publish "douyin/kitex_gen/publish"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishCount implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishCount(ctx context.Context, req *publish.PublishCountRequest) (resp *publish.PublishCountResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishExist implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishExist(ctx context.Context, req *publish.PublishExistRequest) (resp *publish.PublishExistResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishInfo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishInfo(ctx context.Context, req *publish.PublishInfoRequest) (resp *publish.PublishInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishVideoList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishVideoList(ctx context.Context, req *publish.PublishVideoListRequest) (resp *publish.PublishVideoListResponse, err error) {
	// TODO: Your code here...
	return
}
