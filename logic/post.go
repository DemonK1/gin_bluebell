package logic

import (
	"gin_bluebell/dao/mysql"
	"gin_bluebell/models"
	"gin_bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) error {
	// 1. 生成 post id
	p.ID = snowflake.GenID()
	// 2. 保存进数据库
	return mysql.CreatePost(p)
}
