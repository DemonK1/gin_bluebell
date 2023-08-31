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

	v1 := r.Group("/api/v1")

	// 注册路由
	v1.POST("/signup", controllers.SignUpHandler)
	// 登录路由
	v1.POST("/login", controllers.LoginHandler)

	// 应用JWT中间件 与顺序有关 在需要token验证的路由之前注册中间件
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandel)
		v1.GET("/community/:id", controllers.CommunityDetailHandel)
		v1.POST("/post", controllers.CreatePostHandel)
		v1.GET("/post/:id", controllers.GetPostDetailHandel)
		v1.GET("/postList", controllers.GetPostListHandel)
		// 根据时间或分数获取帖子列表
		v1.GET("/postList2", controllers.GetPostListHandel2)
		v1.POST("/vote", controllers.PostVoteHandel)
	}

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
