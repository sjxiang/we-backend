package middleware

import (
	"strings"
	"time"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


// 跨域资源共享
func (h *middleware) Cors() gin.HandlerFunc {
	
	cfg := cors.DefaultConfig()

	cfg.AllowMethods     = []string{"GET", "POST"}
	cfg.AllowHeaders     = []string{"*"}  // e.g. "Authorization"
	cfg.ExposeHeaders    = []string{"*"}
	cfg.AllowCredentials = true
	cfg.MaxAge           = 12 * time.Hour
	
	// 写法 1
	// cfg.AllowOrigins = []string{"http://localhost:8080", "https://www.yourcompany.com"}
	
	// 写法 2
	cfg.AllowOriginFunc = func(origin string) bool {
		if strings.HasPrefix(origin, "http://localhost") {
			// 你的开发环境
			return true
		}
		if strings.Contains(origin, "yourcompany.com") {
			// 你的生产环境
			return true
		}

		return false
	}

	return cors.New(cfg)
}


/*

例：误触微博首页广告，跳转淘宝

跨域资源共享，cors

什么是同源？
	域名、端口、和协议均相同

为什么需要（浏览器）同源策略？
	限制未经允许的访问

怎么实现跨域？
	如果有必要的话，浏览器会发送 preflight，也就是 options 请求

	设置 Header 


*/

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 哪些请求来源是允许的（e.g. 微博首页跳转）
		c.Header("Access-Control-Allow-Origin", "*")
		// 是否允许带上 cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		// 请求可以带上的 Headers
		c.Header("Access-Control-Allow-Headers", "*")
		// 响应返回的非标准 Headers
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
			"Access-Control-Allow-Headers, Authorization, Cache-Control, Content-Language, Content-Type, X-token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Content-Type", "application/json")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}



func (h *middleware) CorsV1() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果 HTTP 请求不是 options 跨域请求，则继续处理 HTTP 请求
		if c.Request.Method != "OPTIONS" {
			c.Next()
		
		// 如果 HTTP 请求是 options 跨域请求，则设置跨域 Header，并返回
		} else {

			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Authorization")
			c.Header("Access-Control-Expose-Headers", "x-token")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
			c.Header("Content-Type", "application/json")
					
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	}
}