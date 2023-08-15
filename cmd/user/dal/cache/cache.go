package cache

import (
	"context"
	"strconv"
)

type (
	User struct{}
)

func (f User) AddUser(ctx context.Context, userId int64) {
	add(rdbFavorites, ctx, strconv.FormatInt(userId, 10)+userSuffix, 1)
}

func (f User) DelUser(ctx context.Context, userId int64) {
	del(rdbFavorites, ctx, strconv.FormatInt(userId, 10)+userSuffix, 1)
}

func (f User) CheckUser(ctx context.Context, userId int64) bool {
	return exist(rdbFavorites, ctx, strconv.FormatInt(userId, 10)+userSuffix, 1)
}
