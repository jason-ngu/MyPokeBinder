package utilities

import (
	"math"
	"net/http"
	"strconv"
)

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

func GetPaginationFromRequest(r *http.Request) *PaginationQuery {
	query := r.URL.Query()
	sizeStr := query.Get("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return nil
	}
	pageStr := query.Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil
	}
	return &PaginationQuery{
		Size: size,
		Page: page,
	}
}
