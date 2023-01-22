package errcode

var (
	ErrorGetTagListFailed = NewError(2001_0001, "获取标签列表失败")
	ErrorCreateTagFailed  = NewError(2001_0002, "创建标签失败")
	ErrorUpdateTagFailed  = NewError(2001_0003, "更新标签失败")
	ErrorDeleteTagFailed  = NewError(2001_0004, "删除标签失败")
	ErrorCountTagFailed   = NewError(2001_0005, "统计标签失败")
)
