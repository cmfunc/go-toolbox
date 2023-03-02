package cacheredis

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-errors/errors"
	"github.com/go-redis/redis/v8"
)

// 基于Redis做缓存时,分页查询场景下的缓存查询与同步封装

type own struct {
	cli             *redis.Client //redis连接客户端
	cacheKey_Prefix string        //缓存前缀
	expire          time.Duration //超时时间
}

type Piker interface{}

func NewOwn(cli *redis.Client, prefix string, ex time.Duration) Piker {
	return &own{
		cli:             cli,
		cacheKey_Prefix: prefix + "::%s",
		expire:          ex,
	}
}

// Insert
// 插入数据时,数据更新时,数据删除时,需要将数据删除
// key: redis缓存时,与prefix格式化拼接成唯一的列表数据标识key,所有的分页数据保存在该key下
// 如果是用户的数据
func (o *own) Del(ctx context.Context, key string) (n int64, err error) {
	key = fmt.Sprint(o.cacheKey_Prefix, key)
	return o.cli.Del(ctx, key).Result()
}

// Query
// 查询数据时,数据库中没有数据,主动加载数据到
func (o *own) Set(ctx context.Context, key string, rows []interface{}, offset uint64) (err error) {
	if len(rows) < 1 {
		return errors.New("rows length is zero")
	}
	zmembers := make([]*redis.Z, 0)
	for i := 0; i < len(rows); i++ {
		rowBytes, err := json.Marshal(rows[i])
		if err != nil {
			return err
		}
		zmembers = append(zmembers, &redis.Z{
			Score:  float64(offset + uint64(i)),
			Member: string(rowBytes),
		})
	}
	key = fmt.Sprint(o.cacheKey_Prefix, key)
	_, err = o.cli.ZAdd(ctx, key, zmembers...).Result()
	if err != nil {
		return err
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

// Get 
// 从缓存中获取分页数据
func (o *own) Get(ctx context.Context, key string, offset, limit uint64) (bytesArray [][]byte, err error) {
	opts := &redis.ZRangeBy{
		Min:    strconv.FormatInt(int64(offset), 10),
		Max:    strconv.FormatInt(int64(offset+limit), 10),
		Offset: int64(offset),
		Count:  int64(limit),
	}
	key = fmt.Sprint(o.cacheKey_Prefix, key)
	strs, err := o.cli.ZRangeByScore(ctx, key, opts).Result()
	if err != nil {
		return nil, err
	}
	bytesArray = [][]byte{}
	for _, str := range strs {
		bytesArray = append(bytesArray, []byte(str))
	}
	return bytesArray, nil
}
