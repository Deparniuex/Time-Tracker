package api

import (
	"example.com/tracker/internal/entity"
	"example.com/tracker/pkg/util"
)

type CreateUserRequest struct {
	PassportNumber string `json:"passportNumber" binding:"required" example:"1234 567890"`
}

type GetUsersRequest struct {
	Pagination
	Name       string `form:"name" binding:"omitempty" example:"Ivan"`
	Surname    string `form:"surname" binding:"omitempty" example:"Ivanov"`
	Patronymic string `form:"patronymic" binding:"omitempty" example:"Ivanovich"`
	Address    string `form:"address" binding:"omitempty" example:"Nizhny Novgorod, Gorky Street 6"`
}

type UpdateUserRequest struct {
	Name           string `json:"name" binding:"required" example:"Ivan"`
	Surname        string `json:"surname" binding:"required" example:"Ivanov"`
	Patronymic     string `json:"patronymic" binding:"required" example:"Ivanovich"`
	Address        string `json:"address" binding:"required" example:"Nizhny Novgorod, Gorky Street 6"`
	PassportSerie  int    `json:"passport_serie" binding:"required" example:"1234"`
	PassportNumber int    `json:"passport_number" binding:"required" example:"567890"`
}

type GetUsersResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Body    []*entity.User `json:"body"`
	Meta    util.Metadata  `json:"meta"`
}
