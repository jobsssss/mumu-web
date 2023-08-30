package main

import (
	"embed"
	middlewares2 "mumu/app/middlewares"
	"mumu/bootstrap"
	"mumu/config"
	c "mumu/pkg/config"
	"mumu/pkg/logger"
	"net/http"
)

//go:embed resources/views/*
var tplFS embed.FS

//go:embed public/*
var staticFS embed.FS

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 初始化 Logger
	bootstrap.SetupLogger()
	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化 Redis
	bootstrap.SetupRedis()

	// 初始化缓存
	bootstrap.SetupCache()

	bootstrap.SetupTemplate(tplFS)

	// 初始化路由绑定
	router := bootstrap.SetupRoute(staticFS)

	err := http.ListenAndServe(":"+c.GetString("app.port"), middlewares2.RemoveTrailingSlash(router))

	logger.LogIf(err)
}
