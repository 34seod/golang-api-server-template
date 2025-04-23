package routes

import (
	"golang-api-server-template/internal/controllers"

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
		v1.GET("/users", controllers.UserFindAll)
		v1.GET("/users/:id", controllers.UserFindByID)
		v1.POST("/users", controllers.UserCreate)
		v1.PUT("/users/:id", controllers.UserUpdate)
		v1.DELETE("/users/:id", controllers.UserDelete)
	}
}
