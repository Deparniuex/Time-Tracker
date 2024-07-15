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

// @Summary Create new user
// @Tags Users
// @Accept json
// @Produce json
// @Param data body api.CreateUserRequest true "Request body"
//
// @Success 201 {object} api.DefaultResponse "User succesfully created"
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /users/create [post]
func (h *Handler) createUser(ctx *gin.Context) {
	logrus.Debug("received createUser request")
	var req api.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	logrus.Debugf("spliting passport numbers: %s", req.PassportNumber)
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

	logrus.Debug("user successfully created")
	ctx.Header("Locations", fmt.Sprintf("/users/%d", user.ID))
	ctx.JSON(http.StatusCreated, &api.DefaultResponse{
		Code:    http.StatusOK,
		Message: "user succesfully created",
	})
}

// @Summary Get users using pagination
// @Tags Users
// @Produce json
// @Param filter query api.GetUsersRequest true "Pagination and filters"
//
// @Success 200 {object} api.GetUsersResponse
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /users/ [get]
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

	logrus.WithFields(logrus.Fields{
		"name":       req.Name,
		"surname":    req.Surname,
		"patronymic": req.Patronymic,
		"address":    req.Address,
	}).Debug("recieved fields")

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

// @Summary      Update user by id.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param data body api.UpdateUserRequest true "Request body"
//
// @Success      200 {object} api.DefaultResponse "User succesfully updated"
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users/update/{id} [put]
func (h *Handler) updateUser(ctx *gin.Context) {
	var reqID api.ID
	err := ctx.ShouldBindUri(&reqID)

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var req api.UpdateUserRequest
	err = ctx.ShouldBindJSON(&req)

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = h.Service.UpdateUser(&entity.User{
		ID:             reqID.Value,
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

// @Summary      Delete user by ID
// @Tags         Users
// @Produce      json
// @Param        id   path      int  true  "User ID"
//
// @Success      200 {object} api.DefaultResponse
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users/delete/{id} [delete]
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

	logrus.WithField("userID", id).Debug("Received ID")
	err = h.Service.DeleteUser(id.Value)

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
		Message: "user succesfully deleted",
	})
}
