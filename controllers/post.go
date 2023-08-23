package controllers

import (
	"fmt"
	"gin_bluebell/logic"
	"gin_bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandel 新增帖子
func CreatePostHandel(c *gin.Context) {
	// 1. 参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 从c中取到当前发请求的用户的id值
	userID, err := getCurrentUserID(c)
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

// GetPostDetailHandel 查询帖子详情
func GetPostDetailHandel(c *gin.Context) {
	// 1. 获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostDetailHandel failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 获取帖子详情
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回参数
	ResponseSuccess(c, data)
}

// GetPostListHandel 获取帖子列表
func GetPostListHandel(c *gin.Context) {
	// 获取分页参数
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}

	// 获取数据
	list, err := logic.GetPostList(page, size)
	fmt.Printf("我是错误%v\n", err)
	if err != nil {
		return
	}
	// 返回相应
	ResponseSuccess(c, list)
}
