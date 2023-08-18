package controllers

import (
	"fmt"
	"gin_bluebell/logic"
	"gin_bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误,直接返回
		zap.L().Error("ShouldBindJSON with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(
				http.StatusOK, gin.H{
					"msg": err.Error(),
				},
			)
			return
		}
		c.JSON(
			http.StatusOK, gin.H{
				"msg": errs.Translate(trans),
			},
		)
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
		fmt.Printf("注册失败错误: %v", err)
		c.JSON(
			http.StatusOK, gin.H{
				"msg": "注册失败",
			},
		)
		return
	}
	// 3.返回响应
	c.JSON(
		http.StatusOK, gin.H{
			"msg": "success",
		},
	)
}
