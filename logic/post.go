package logic

import (
	"gin_bluebell/dao/mysql"
	"gin_bluebell/models"
	"gin_bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	// 1. 生成 post id
	p.ID = snowflake.GenID()
	// 2. 保存进数据库
	return mysql.CreatePost(p)
}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询并拼接接口想用的数据
	// 查询帖子信息
	postData, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed", zap.Error(err))
		return
	}
	// 根据用户id查询用户详情信息
	userData, err := mysql.GetUserById(postData.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(postData.AuthorID) failed", zap.Error(err))
		return
	}
	// 根据社区id查询社区详情id
	communityData, err := mysql.GetCommunityDetailByID(postData.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(postData.CommunityID)", zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      userData.Username,
		Post:            postData,
		CommunityDetail: communityData,
	}
	return
}
