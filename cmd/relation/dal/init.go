package dal

import (
	"douyin/cmd/relation/dal/cache"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/dal/mq"
)

func Init() {
	db.Init()
	cache.Init()
	mq.Init()
}
