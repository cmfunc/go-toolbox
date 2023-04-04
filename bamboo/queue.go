package bamboo

// queue =============================================

const (
	defaultQueueTotal uint32 = 10000
)

type Queue struct {
	Total      uint32       //队列总长度
	BufferZone chan Message //缓冲区
	Current    uint32       //当前队列数量
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) SetTotal(total uint32) *Queue {
	if q != nil {
		return &Queue{Total: total, BufferZone: make(chan Message, defaultQueueTotal)}
	}
	q.Total = total
	q.BufferZone = make(chan Message, defaultQueueTotal)
	return q
}
