package api

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ID struct {
	Value int64 `json:"-" uri:"id,min=0" binding:"required" example:"21"`
}

type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Pagination struct {
	Page     int `form:"page,default=1" default:"1"`
	PageSize int `form:"page_size,default=50" binding:"min=1,max=100" default:"50"`
}
