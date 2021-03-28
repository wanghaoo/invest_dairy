package util

import "gorm.io/gorm"

type Pagination struct {
	PageNo   int
	PageSize int
}

func (p *Pagination) Start() int {
	if p.PageNo > 0 && p.PageSize > 0 {
		page := p.PageNo
		perPage := p.PageSize
		return (page - 1) * perPage
	}
	return 0
}

func (p *Pagination) PageLimit(q *gorm.DB) *gorm.DB {
	if p.PageNo > 0 && p.PageSize > 0 {
		return q.Offset(p.Start()).Limit(p.PageSize)
	}
	return q
}

type ResponseData struct {
	Data interface{} `json:"data"`
	Page Page        `json:"page"`
	Err  error       `json:"err"`
}

type (
	Page struct {
		PageNo   uint64 `json:"pageNo" desc:"第几页(1开始计数)"`
		PageSize uint64 `json:"pageSize" desc:"每页有几条"`
		RowsNo   uint64 `json:"rowsNo" desc:"总条数"`
		PagesNo  uint64 `json:"pagesNo" desc:"总页数"`
	}
)
