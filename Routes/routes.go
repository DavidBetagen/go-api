package Routes

import (
	"go-api/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	user := router.Group("/user")
	{
		UserController := new(Controllers.UserController)
		user.GET("user", UserController.GetUsers)
		user.POST("user", UserController.CreateUser)
		user.GET("user/:id", UserController.GetUserByID)
		user.PUT("user/:id", UserController.UpdateUser)
		user.DELETE("user/:id", UserController.DeleteUser)
	}
	return router
}
