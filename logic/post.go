package logic

import (
	"gin_bluebell/dao/mysql"
	"gin_bluebell/dao/redis"
	"gin_bluebell/models"
	"gin_bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成 post id
	p.ID = snowflake.GenID()
	// 2. 保存进数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID)
	return
}

// GetPostById 获取帖子详情
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

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	postList, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(postList))
	for _, post := range postList {
		// 根据用户id获取用户详情
		userData, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(postData.AuthorID) failed", zap.Error(err))
			continue
		}
		// 根据社区id获取社区详情
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(postData.CommunityID)", zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      userData.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

// GetPostList2 根据时间或分数获取帖子列表
func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 1. 去redis查询id列表
	ids, err := redis.GetPostIDInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		return
	}
	// 2. 根据id去mysql数据库查询帖子详情信息
	postList, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	// 将帖子的作者及分区信息查询出来 填充到帖子中
	for _, post := range postList {
		// 根据用户id获取用户详情
		userData, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(postData.AuthorID) failed", zap.Error(err))
			continue
		}
		// 根据社区id获取社区详情
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(postData.CommunityID)", zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      userData.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
