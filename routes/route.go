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
	r.POST("/signup", controllers.SignUpHandler)

	r.GET(
		"/", func(c *gin.Context) {
			c.String(http.StatusOK, "ok!")
		},
	)
	if err := r.Run("127.0.0.1:8080"); err != nil {
		panic(err)
	}
	return r
}
