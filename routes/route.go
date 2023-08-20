package routes

// route 路由的配置及初始化

import (
	"fmt"
	"gin_bluebell/controllers"
	"gin_bluebell/logger"
	"gin_bluebell/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}

	// 初始化gin框架内置的校验器使用的翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed,err:%v\n", err)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册路由
	r.POST("/signup", controllers.SignUpHandler)
	// 登录路由
	r.POST("/login", controllers.LoginHandler)
	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 如果是登陆用户判断请求头中是否有 有效的JWT
		c.String(http.StatusOK, "pong")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	if err := r.Run("127.0.0.1:8080"); err != nil {
		panic(err)
	}
	return r
}
