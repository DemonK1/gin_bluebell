package mysql

import (
	"gin_bluebell/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
    post_id,title,content,author_id,community_id)
    values(?,?,?,?,?)
  `
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostById(pid int64) (postData *models.Post, err error) {
	postData = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post where post_id=?`
	err = db.Get(postData, sqlStr, pid)
	return
}
