package handler

import (
	"time"

	"example.com/tracker/internal/service"
)

type HandlerConfig struct {
	Port    string
	Host    string
	Timeout time.Duration
}

type Handler struct {
	Config  HandlerConfig
	Service service.Service
}
