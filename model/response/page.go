package response

type PageResponse struct {
	Total   int64       `json:"total"`   // 总记录数
	Size    int         `json:"size"`    // 每页的记录数
	Pages   int         `json:"pages"`   // 总页数
	Records interface{} `json:"records"` // 分页返回的具体记录
	Current int         `json:"current"` // 当前页码
}
