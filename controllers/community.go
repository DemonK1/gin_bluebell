package controllers

import (
	"gin_bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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
