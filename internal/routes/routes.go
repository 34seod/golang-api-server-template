package routes

import (
	"golang-api-server-template/internal/controller"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// v1
	{
		v1 := r.Group("/v1")
		v1.GET("/users", controller.UserFindAll)
		v1.GET("/users/:id", controller.UserFindByID)
		v1.POST("/users", controller.UserCreate)
		v1.PUT("/users/:id", controller.UserUpdate)
		v1.DELETE("/users/:id", controller.UserDelete)
	}
}
