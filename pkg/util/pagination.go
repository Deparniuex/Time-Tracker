package util

import "math"

type Pagination struct {
	Page     int
	PageSize int
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `jEmptyson:"total_records,omitempty"`
}

func New(page int, pageSize int) *Pagination {
	return &Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}

func (f *Pagination) Offset() int {
	return (f.Page - 1) * f.PageSize
}

func (f *Pagination) CalculateMetadata(totalRecords int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  f.Page,
		PageSize:     f.PageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(f.PageSize))),
		TotalRecords: totalRecords,
	}
}
