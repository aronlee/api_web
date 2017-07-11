package models

// Page 分页
type Page struct {
	PageNo     int         // 页数
	PageSize   int         // 每页个数
	TotalPage  int         // 总页数
	TotalCount int         // 总数
	FirstPage  bool        // 是否是首页
	LastPage   bool        // 是否是最后一页
	List       interface{} // 列表
}

// PageUtil 分页工具类
func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	var tp int
	if pageSize > 0 {
		tp = count / pageSize
		if count%pageSize > 0 {
			tp = count/pageSize + 1
		}
	}
	return Page{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalPage:  tp,
		TotalCount: count,
		FirstPage:  pageNo == 1,
		LastPage:   pageNo == tp,
		List:       list,
	}
}
