package test

import (
	"github.com/cxp-13/c2n-launchpad-go-be/internal/portalapi/handlers"
	"github.com/cxp-13/c2n-launchpad-go-be/internal/portalapi/service"
	"github.com/gin-gonic/gin"
)

func SetupGinRouter() *gin.Engine {
	r := gin.Default()

	userService := service.NewUserService("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	userHandler := handlers.NewUserHandler(userService)

	projectService := &service.ProjectService{}
	productHandler := handlers.NewProductHandler(projectService)

	r.POST("/registered", userHandler.Registered)
	r.POST("/participation", userHandler.Participation)
	productRoutes := r.Group("/product")
	{
		productRoutes.GET("/base_info", productHandler.BaseInfo)
		productRoutes.GET("/list", productHandler.List)
	}

	return r
}
