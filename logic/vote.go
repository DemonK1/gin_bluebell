package logic

import (
	"gin_bluebell/dao/redis"
	"gin_bluebell/models"
	"strconv"
)

// VoteForPost 为帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
