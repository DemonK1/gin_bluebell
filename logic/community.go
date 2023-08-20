package logic

import (
	"gin_bluebell/dao/mysql"
	"gin_bluebell/models"
)

// GetCommunityList 查询列表数据
func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查到所有的community并返回
	return mysql.GetCommunityList()
}

// GetCommunityDetailList 查询列表详情数据
func GetCommunityDetailList(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
