package oncecache

import (
	"context"
	"sync"
	"time"
)

var once = &sync.Once{}
var ccond *cases

func Set(ctx context.Context, key string, val interface{}, ex time.Duration) error {
	return ccond.Set(ctx, key, val, ex)
}
func Get(ctx context.Context, key string) (val interface{}, err error) {
	return ccond.Get(ctx, key)
}

type kvItem struct {
	val interface{}
	ex  time.Duration
}

type cases struct {
	kve   map[string]kvItem
	mutex *sync.RWMutex
}

func init() {
	once.Do(func() { ccond = NewCaseMap() })
}

func NewCaseMap() *cases {
	return &cases{
		kve:   map[string]kvItem{},
		mutex: &sync.RWMutex{},
	}
}

// key:指定标识
// val:建议为函数或通用接口
func (c *cases) Set(ctx context.Context, key string, val interface{}, ex time.Duration) (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.kve[key] = kvItem{
		val: val,
		ex:  ex,
	}
	return nil
}

func (c *cases) Get(ctx context.Context, key string) (val interface{}, err error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.kve[key], nil
}
