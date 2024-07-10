package handler

import (
	"example.com/tracker/internal/service"
)

type Handler struct {
	Service service.Service
}

func New(srvc service.Service) *Handler {
	return &Handler{
		Service: srvc,
	}
}
