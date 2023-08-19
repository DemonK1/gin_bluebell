package controllers

import (
	"errors"
	"gin_bluebell/dao/mysql"
	"gin_bluebell/logic"
	"gin_bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 注册函数
func SignUpHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误,直接返回
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}
	// 手动判断参数
	// if len(p.Username) == 0 || len(p.Password) == 0 || len(p.Repassword) == 0 || p.Repassword != p.Password {
	//   // 请求参数有误,直接返回
	//   zap.L().Error("ShouldBindJSON with invalid param")
	//   c.JSON(
	//     http.StatusOK, gin.H{
	//       "msg": "请求参数有误",
	//     },
	//   )
	//   return
	// }

	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登录函数
func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	// 1.参数校验
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误,直接返回
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}
	// 2.业务逻辑处理
	if err := logic.Login(p); err != nil {
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}
