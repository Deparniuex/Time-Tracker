package api

import "example.com/tracker/internal/entity"

type GetUsersRequest struct {
	Name       string `json:"name" binding:"omitempty"`
	Surname    string `json:"surname" binding:"omitempty"`
	Patronymic string `json:"patronymic" binding:"omitempty"`
	Address    string `json:"address" binding:"omitempty"`
}

type GetUsersResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Body    []*entity.User `json:"body"`
}
