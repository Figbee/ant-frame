package router

import "github.com/gin-gonic/gin"

func InitPublicRouter(g *gin.RouterGroup) (R gin.IRoutes) {
	g.POST("/login")
	return g
}
