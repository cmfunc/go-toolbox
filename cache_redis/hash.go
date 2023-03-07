package cacheredis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

// 使用hash结构缓存
type hash struct {
	cli             *redis.Client //redis连接客户端
	cacheKey_Prefix string        //缓存前缀
	expire          time.Duration //超时时间
}

func NewHash(cli *redis.Client, prefix string, ex time.Duration) *hash {
	return &hash{
		cli:             cli,
		cacheKey_Prefix: prefix + "::%s",
		expire:          ex,
	}
}

