package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-example/controller"
)

func userRouter(user *gin.RouterGroup)  {
	user.GET("/:id",controller.GetUser)
	user.GET("/",controller.GetAllUser)
	user.POST("/create",controller.CreateUser)
}
