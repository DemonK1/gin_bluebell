package routes

// route 路由的配置及初始化

import (
	"gin_bluebell/controllers"
	"gin_bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册路由
	r.GET("/signup", controllers.SignUpHandler)

	r.GET(
		"/", func(c *gin.Context) {
			c.String(http.StatusOK, "ok!")
		},
	)
	return r
}
