package util

import "github.com/Masterminds/squirrel"

type Pager struct {
  PageNo   uint64 `json:"pageNo" desc:"第几页(1开始计数)"`
  PageSize uint64 `json:"pageSize" desc:"每页有几条"`
}

func (p *Pager) Build() {
  if p.PageNo <= 0 {
    p.PageNo = 1
  }
  if p.PageSize <= 0 {
    p.PageSize = 10
  }
}

func (p *Pager) Start() uint64 {
  if p.PageNo > 0 && p.PageSize > 0 {
    page := p.PageNo
    perPage := p.PageSize
    return (page - 1) * perPage
  }
  return 0
}

func (p *Pager) PageLimit(q squirrel.SelectBuilder) squirrel.SelectBuilder {
  if p.PageSize == 0 {
    p.PageSize = 10
  }
  if p.PageNo == 0 {
    p.PageNo = 1
  }
  return q.Offset(p.Start()).Limit(p.PageSize)
}
