package mysql

import (
	"gin_bluebell/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
    post_id,title,content,author_id,community_id)
    values(?,?,?,?,?)
  `
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostById 获取帖子详情
func GetPostById(pid int64) (postData *models.Post, err error) {
	postData = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post where post_id=?`
	err = db.Get(postData, sqlStr, pid)
	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (postList []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post limit ?,?`
	postList = make([]*models.Post, 0, 2)
	err = db.Select(&postList, sqlStr, (page-1)*size, size)
	return
}

// GetPostListByIDs 根据给定的id查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id,?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
