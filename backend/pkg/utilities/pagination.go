package utilities

import "math"

const defaultSize = 10

type PaginationQuery struct {
	Size int `json:"size"`
	Page int `json:"page"`
}

func NewPaginationQuery(size int, page int) *PaginationQuery {
	return &PaginationQuery{
		Size: size,
		Page: page,
	}
}

func (p *PaginationQuery) GetSize() int {
	return p.Size
}

func (p *PaginationQuery) GetPage() int {
	return p.Page
}

func (p *PaginationQuery) GetOffset() int {
	if p.Page == 0 {
		return 0
	}
	return (p.Page - 1) * p.Size
}

func (p *PaginationQuery) GetLimit() int {
	return p.Size
}

// Get total pages int
func GetTotalPages(totalCount int, pageSize int) int {
	d := float64(totalCount) / float64(pageSize)
	return int(math.Ceil(d))
}
