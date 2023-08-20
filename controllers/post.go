package controllers

import (
	"fmt"
	"gin_bluebell/logic"
	"gin_bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandel(c *gin.Context) {
	// 1. 参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 从c中取到当前发请求的用户的id值
	userID, err := getCurrentUserID(c)
	fmt.Printf("111111111:%v\n", err)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}
