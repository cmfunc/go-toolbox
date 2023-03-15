package oncecache

import (
	"errors"
	"sync"
	"time"
)

// 定时器池化技术
type timerLoop struct {
	tMin    uint32        //最大定时器个数
	tMax    uint32        //最少定时器个数
	wait    []*time.Timer //等待执行
	current uint64        //当前timer总数
	idle    []*time.Timer //空闲timer
	mux     *sync.Mutex   //运行时锁
}

type Option func(*timerLoop)

func SetMax(max uint32) Option {
	return func(loop *timerLoop) {
		loop.tMax = max
	}
}

// 构造函数
func NewTimerLoop(opts ...Option) *timerLoop {
	loop := &timerLoop{}
	for _, opt := range opts {
		opt(loop)
	}
	return loop
}

// 添加任务时,先判断当前可用的定时器个数
// 定时器不够用时,不申请新的goroutine,直接阻塞或返回定时器数量不够的错误

func (loop *timerLoop) addTask(timer *time.Timer, ex time.Duration, do func()) (err error) {
	if timer == nil {
		return errors.New("timer uninitial")
	}
	timer.Reset(ex)

	<-timer.C
	do()
	loop.idle = append(loop.idle, timer)

	return nil
}
