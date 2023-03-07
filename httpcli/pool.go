package httpcli

import "sync"

var cliPool = sync.Pool{
	New: func() any {
		return newClient()
	},
}

func New() *Client {
	cli := cliPool.Get().(*Client)
	return cli
}

// Close
// 关闭body
// 回收http client
func (c *Client) Close() {
	cliPool.Put(c)
}
