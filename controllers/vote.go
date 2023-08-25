package controllers

import (
	"gin_bluebell/logic"
	"gin_bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// PostVoteHandel 投票
func PostVoteHandel(c *gin.Context) {
	// 1. 参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) // 翻译并去除掉提示中的结构体标
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	// 2. 具体的投票逻辑处理
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost(userID, p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}
