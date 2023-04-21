package bamboo

import (
	"context"
	"sync/atomic"
)

func Publish(ctx context.Context, topic string, msg Message) error {
	queue, ok := broker.MQ[Topic(topic)]
	if !ok {
		broker.MQ[Topic(topic)] = Queue{Total: defaultQueueTotal, BufferZone: make(chan Message, defaultQueueTotal)}
	}
	queue.BufferZone <- msg
	atomic.AddUint32(&queue.Current, 1)
	return nil
}
