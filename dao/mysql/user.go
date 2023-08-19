package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"gin_bluebell/models"
)

const secret = "mainarr@yeah.net"

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

func InsertUser(user *models.User) (err error) {
	// 密码加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	pswMd5 := h.Sum([]byte(oPassword))
	return hex.EncodeToString(pswMd5)
}

func Login(p *models.ParamLogin) (err error) {
	var user models.User
	oPassword := p.Password // 用户输入的密码
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(&user, sqlStr, p.Username)
	// 判断用户是否存在
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	// 查询数据库中是否存在
	if err != nil {
		return
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return errors.New("密码错误")
	}
	return nil
}
