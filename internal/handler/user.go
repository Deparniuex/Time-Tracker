package handler

import (
	"fmt"
	"net/http"

	"example.com/tracker/internal/entity"
	"example.com/tracker/internal/handler/api"
	"example.com/tracker/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var req api.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	passNumber, passSerie, err := util.SplitPassport(req.PassportNumber)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	user := &entity.User{
		PassportSerie:  passSerie,
		PassportNumber: passNumber,
	}
	err = h.Service.CreateUser(user)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	ctx.Header("Locations", fmt.Sprintf("/users/%d", user.ID))
	ctx.JSON(http.StatusOK, &api.DefaultResponse{
		Code:    http.StatusOK,
		Message: "user succesfully created",
	})
}

func (h *Handler) getUsers(ctx *gin.Context) {
	var req api.GetUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	logrus.Info(req.Name, req.Surname, req.Patronymic, req.Address)
	filters := map[string]string{
		"first_name": req.Name,
		"surname":    req.Surname,
		"patronymic": req.Patronymic,
		"address":    req.Address,
	}
	users, metadata, err := h.Service.GetUsers(util.New(req.Page, req.PageSize), filters)

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.GetUsersResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Body:    users,
		Meta:    *metadata,
	})
}

func (h *Handler) updateUser(ctx *gin.Context) {
	var req api.UpdateUserRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	logrus.Info(req.ID)
	err = h.Service.UpdateUser(&entity.User{
		ID:             req.ID.Value,
		First_name:     req.Name,
		Surname:        req.Surname,
		Patronymic:     req.Patronymic,
		Address:        req.Address,
		PassportSerie:  req.PassportNumber,
		PassportNumber: req.PassportSerie,
	})
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.DefaultResponse{
		Code:    http.StatusOK,
		Message: "user succesfully updated",
	})
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	var id api.ID
	err := ctx.ShouldBindUri(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	logrus.WithField("userID", id).Info("Received ID")
	err = h.Service.DeleteUser(id.Value)

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.GetUsersResponse{
		Code:    http.StatusOK,
		Message: "user succesfully deleted",
	})
}
