package offset

import "context"

// TransOffset
// 接口列表分页与数据库分页查询参数转换
// @param index 页码
// @param size 页面数据个数
func TransOffset(ctx context.Context, index, size uint64) (offset, limit uint64) {
	offset, limit = 0, 10
	if size > 0 {
		limit = size
	}
	if index > 0 {
		offset = (index - 1) * limit
	}
	return offset, limit
}
