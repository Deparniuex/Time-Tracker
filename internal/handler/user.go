package handler

import (
	"net/http"

	"example.com/tracker/internal/handler/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) getUsers(ctx *gin.Context) {
	var req api.GetUsersRequest
	if err := ctx.ShouldBindQuery(req); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	users, err := h.Service.GetUsers(ctx)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, &api.GetUsersResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Body:    users,
	})
}
