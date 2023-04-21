package bamboo

import (
	"context"
	"sync/atomic"

	"github.com/pkg/errors"
)

// broker ===============================================
var broker = &Broker{
	MaxGo: defaultWorkerNum,
	MQ:    map[Topic]Queue{},
}

const defaultWorkerNum = 100

type Broker struct {
	MaxGo int32           //最大并发数
	MQ    map[Topic]Queue //MQ数据
}

func NewBroker() *Broker {
	return &Broker{}
}

func (b *Broker) SetMaxGo(max int32) *Broker {
	if b == nil {
		return &Broker{MaxGo: max}
	}
	b.MaxGo = max
	return b
}

func (b *Broker) InitTopic(topic string, q Queue) *Broker {
	if b == nil {
		return &Broker{MaxGo: defaultWorkerNum, MQ: map[Topic]Queue{Topic(topic): q}}
	}
	if q.BufferZone == nil {
		panic("q's buffer zone is nil")
	}
	b.MQ[Topic(topic)] = q
	return b
}

func (b *Broker) Publish(ctx context.Context, topic string, msg Message) error {
	queue, ok := b.MQ[Topic(topic)]
	if !ok {
		b.MQ[Topic(topic)] = Queue{Total: defaultQueueTotal, BufferZone: make(chan Message, defaultQueueTotal)}
	}
	queue.BufferZone <- msg
	atomic.AddUint32(&queue.Current, 1)
	return nil
}

func (b *Broker) Subscribe(ctx context.Context, topic string, handle func(ctx context.Context, msg Message) error) (err error) {
	queue, ok := b.MQ[Topic(topic)]
	if !ok {
		return errors.Errorf("unknown topic:[%s]", topic)
	}
	msg := <-queue.BufferZone
	defer func() {
		if err != nil {
			queue.BufferZone <- msg
			queue.Current++
			return
		}
		queue.Current--
		one1 := int32(-1)
		one2 := uint32(one1)
		atomic.AddUint32(&queue.Current, one2)
	}()
	// 消费
	return handle(ctx, msg)
}
