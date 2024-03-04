package request

// PageRequest 分页请求
type PageRequest struct {
	Current   int    `json:"current"`   // 当前页号
	PageSize  int    `json:"pageSize"`  // 页的大小
	SortField string `json:"sortField"` // 升序、降序
	SortOrder string `json:"sortOrder"` // 排序字段
}
