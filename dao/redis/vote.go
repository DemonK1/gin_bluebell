package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

// 本项目使用简化版的投票分数
// 投一票就加432分 86400/200 -> 200张赞成票可以给你的帖子续一天 -> <出自redis实战这本书>
// 一天86400秒

/*
	投票的几种情况

direction=1时,有两种情况:
 1. 之前没投过票,现在头赞成票
 2. 之前投反对票,现在改投赞成票

direction=0时,有两种情况:
 1. 之前投过赞成票,现在取消投票
 2. 之前投过反对票,现在取消投票

direction=-1时,有两种情况:
 1. 之前没投过票,现在头反对票
 2. 之前投赞成票,现在改投反对票

投票的限制:
每个帖子自发表之日起一个星期之内允许用户投票,超过一个星期就不允许再投票了
 1. 到期之后将redis中保存的赞成票及反对票存储到MySQL表中
 2. 到期之后删除那个 KeyPostVotedZSetPF
*/
const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func CreatePost(postID int64) error {
	// redis事务操作
	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  0,
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

// VoteForPost 投票
func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断投票限制
	// 去redis取帖子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2和3需要放到一个pipeline事务中操作
	// 2. 更新帖子分数
	// 先查当前用户给当前帖子的投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	// 方向
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	// 3. 记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
