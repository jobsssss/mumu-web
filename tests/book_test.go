package tests

import (
	"mumu/app/models/chapter"
	"mumu/bootstrap"
	"testing"
)

func Test_BookModelAjaxChapter(t *testing.T) {
	// 初始化 Logger
	bootstrap.SetupLogger()
	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化 Redis
	bootstrap.SetupRedis()

	// 初始化缓存
	bootstrap.SetupCache()

	chapter.PageFind(354, 1, "ASC")
}
