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
