package redis

// redis 的配置及初始化

import (
	"fmt"
	"gin_bluebell/models"
	"gin_bluebell/settings"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf(
				"%s:%d", cfg.Host, cfg.Port,
			),
			Password: cfg.Password, // no password set
			DB:       cfg.DB,       // use default DB
			PoolSize: cfg.PoolSize,
		},
	)

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}

func GetPostIDInOrder(p *models.ParamPostList) ([]string, error) {
	// 1. 从redis获取id
	// 2. 根据用户请求携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 3. 确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// 4. ZREVRANGE 按分数从大到小的顺序查询指定数量的元素
	return rdb.ZRevRange(key, start, end).Result()
}
