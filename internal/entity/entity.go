package entity

import "regexp"

type ID string
type Email string

// Validate is a function to validate email
func (e *Email) Validate() bool {
	pattern := regexp.MustCompile(`^(\w|\.)+@(\w)+(.(\w)+){1,2}$`)
	return pattern.MatchString(e.String())
}

type ParamOption struct {
	OrderByField string // ex: order by 'name'
	OrderByDir   string // ex: desc || 'asc'
}

type Pagination struct {
	Limit int `json:"limit,omitempty;query:limit"`
	Page  int `json:"page,omitempty;query:page"`
	// Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}
