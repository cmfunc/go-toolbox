package casemap

import (
	"context"
	"sync"
)

/*
使用map解决众多if,和switch case
*/
type cases struct {
	kv    map[string]interface{}
	mutex *sync.RWMutex
}

func NewCaseMap() *cases {
	return &cases{
		kv:    map[string]interface{}{},
		mutex: &sync.RWMutex{},
	}
}

// key:指定标识
// val:建议为函数或通用接口
func (c *cases) Set(ctx context.Context, key string, val interface{}) (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.kv[key] = val
	return nil
}

func (c *cases) Get(ctx context.Context, key string) (val interface{}, err error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.kv[key], nil
}
