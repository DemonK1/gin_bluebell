package logic

import (
	"gin_bluebell/dao/mysql"
	"gin_bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查到所有的community并返回
	return mysql.GetCommunityList()
}
