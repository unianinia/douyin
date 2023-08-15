package dal

import (
	"douyin/cmd/favorite/dal/cache"
	"douyin/cmd/favorite/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
