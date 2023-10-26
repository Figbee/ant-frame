package router

import (
	"ant-frame/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "i am ok!",
		})
	})

	versionGroup := r.Group("/v1")
	//初始化公共路由
	router := InitPublicRouter(versionGroup)
	//初始化用户路由

}
