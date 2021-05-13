package main

import (
	_ "github.com/Kawaii-jump/gin-admin/docs"
	"github.com/Kawaii-jump/gin-admin/router"
)

// @title grpc MQ API
// @version 1.0
// @description This is grpc message queue server api

// @contact.name hetao.kawaii
// @contact.url https://hetao.kawaii.com
// @contact.email hetao.kawaii@bytedance.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8181
// @BasePath /api/v1

func main() {
	router.InitRouter()
}
