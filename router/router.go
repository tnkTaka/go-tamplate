package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter()  *gin.Engine{
	r := gin.Default()

	//r.Static("/js", "./public/js")
	//r.Static("/image", "./public/image")
	//r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")

	//r.NoRoute(func(c *gin.Context) {
	//	c.HTML(200, "index.html", nil)
	//})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	user := r.Group("/user")
	userRouter(user)

	return r
}
