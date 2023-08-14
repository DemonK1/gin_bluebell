package logic

import (
	"gin_bluebell/dao/mysql"
	"gin_bluebell/pkg/snowflake"
)

func SignUp() {
	// 1.判断用户存不存在
	mysql.QueryUserByUsername()
	// 2.生成UID
	snowflake.GenID()
	// 3.密码加密
	// 4.保存进数据库
	mysql.InsertUser()
}
