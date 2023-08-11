package entity

type ID string
type Email string

type ParamOption struct {
	OrderByField string // ex: order by 'name'
	OrderByDir   string // ex: desc || 'asc'
}

type Pagination struct {
	Limit  int
	Offset int
}
