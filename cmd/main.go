package main

import (
	"fmt"
	"github.com/cxp-13/c2n-launchpad-go-be/internal/portalapi/handler"
	"github.com/cxp-13/c2n-launchpad-go-be/internal/portalapi/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 实例化UserService
	userService := service.NewUserService("private_key")
	userHandler := handler.NewUserHandler(userService)

	// 创建 Gin 的默认路由引擎
	r := gin.Default()

	// 注册路由
	r.POST("/registered", userHandler.Registered)
	r.POST("/participation", userHandler.Participation)

	// 启动服务器
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
