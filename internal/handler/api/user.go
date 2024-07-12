package api

import "example.com/tracker/internal/entity"

type CreateUserRequest struct {
	PassportNumber string `json:"passportNumber" binding:"required"`
}

type GetUsersRequest struct {
	Name       string `json:"name" binding:"omitempty"`
	Surname    string `json:"surname" binding:"omitempty"`
	Patronymic string `json:"patronymic" binding:"omitempty"`
	Address    string `json:"address" binding:"omitempty"`
}

type UpdateUserRequest struct {
	ID
	Name           string `json:"name" binding:"omitempty"`
	Surname        string `json:"surname" binding:"omitempty"`
	Patronymic     string `json:"patronymic" binding:"omitempty"`
	Address        string `json:"address" binding:"omitempty"`
	PassportSerie  int    `json:"passport_serie" binding:"omitempty"`
	PassportNumber int    `json:"passport_number" binding:"omitempty"`
}

type GetUsersResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Body    []*entity.User `json:"body"`
}
