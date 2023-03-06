package cacheredis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-errors/errors"
	"github.com/go-redis/redis/v8"
)

/**
基于redis的list数据类型，进行分页查询
list本质底层结构为：
*/

type list struct {
	cli             *redis.Client //redis连接客户端
	cacheKey_Prefix string        //缓存前缀
	expire          time.Duration //超时时间
}

func NewList(cli *redis.Client, prefix string, ex time.Duration) *list {
	return &list{
		cli:             cli,
		cacheKey_Prefix: prefix + "::%s",
		expire:          ex,
	}
}

func (o *list) Del(ctx context.Context, key string) (n int64, err error) {
	key = fmt.Sprint(o.cacheKey_Prefix, key)
	return o.cli.Del(ctx, key).Result()
}
func (o *list) Set(ctx context.Context, key string, rows []interface{}, offset int64) (err error) {
	if len(rows) < 1 {
		return errors.New("rows length is zero")
	}
	key = fmt.Sprint(o.cacheKey_Prefix, key)
	for i := 0; i < len(rows); i++ {
		rowBytes, err := json.Marshal(rows[i])
		if err != nil {
			return err
		}
		_, err = o.cli.LSet(ctx, key, offset+int64(i), string(rowBytes)).Result()
		if err != nil {
			return err
		}
	}
	ok, err := o.cli.Expire(ctx, key, o.expire).Result()
	if err != nil {
		return err
	}
	if !ok {
		return errors.Errorf("Expire exec with key:[%s] ex:[%d] effect item nums:%t", key, o.expire, ok)
	}
	return nil
}

func (o *list) Get(ctx context.Context, key string, offset, limit int64) (bytesArray [][]byte, err error) {
	key = fmt.Sprint(o.cacheKey_Prefix, key)
	strs, err := o.cli.LRange(ctx, key, offset, limit).Result()
	if err != nil {
		return nil, err
	}
	bytesArray = [][]byte{}
	for _, str := range strs {
		bytesArray = append(bytesArray, []byte(str))
	}
	return bytesArray, nil
}
