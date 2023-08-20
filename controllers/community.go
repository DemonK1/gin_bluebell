package controllers

import (
	"gin_bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CommunityHandel 查询社区列表
func CommunityHandel(c *gin.Context) {
	// 查询到所有社区 (community_id,community_name) 以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandel 根据列表id查询详情数据
func CommunityDetailHandel(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("id") // 获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetailList(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetailList(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错信息暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
