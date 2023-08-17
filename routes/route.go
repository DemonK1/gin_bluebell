package routes

// route 路由的配置及初始化

import (
	"fmt"
	"gin_bluebell/controllers"
	"gin_bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	// 初始化gin框架内置的校验器使用的翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed,err:%v\n", err)
	}

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
