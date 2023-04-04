package bamboo

import (
	"context"
	"sync/atomic"

	"github.com/pkg/errors"
)

func Subscribe(ctx context.Context, topic string, handle func(ctx context.Context, msg Message) error) (err error) {
	queue, ok := broker.MQ[Topic(topic)]
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
