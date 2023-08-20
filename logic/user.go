package logic

import (
	"gin_bluebell/dao/mysql"
	"gin_bluebell/models"
	"gin_bluebell/pkg/jwt"
	"gin_bluebell/pkg/snowflake"
)

// SignUp 注册逻辑处理
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	// 2.生成UID
	userID := snowflake.GenID()
	// 构造一个User实例
	user := &models.User{
		UserId:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 4.保存进数据库
	return mysql.InsertUser(user)
}

// Login 登录逻辑处理
func Login(p *models.ParamLogin) (token string, err error) {
	if err := mysql.Login(p); err != nil {
		return "", err
	}
	var user models.User
	token, err = jwt.GenToken(user.UserId, user.Username)
	if err != nil {
		return "", err
	}
	return token, err
}
