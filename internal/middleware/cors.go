package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 允许访问所有域
		c.Header("Access-Control-Allow-Origin", "*")
		// header的类型
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
		// 跨域关键设置 让浏览器可以解析
		c.Header("Access-Control-Allow-Credentials", "true")
		// 跨域请求方法
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		// 服务器支持的所有跨域请求的方法
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		// 跨域请求是否需要带cookie信息 设置true
		if method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()
	}
}
