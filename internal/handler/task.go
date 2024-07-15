package handler

import (
	"fmt"
	"net/http"

	"example.com/tracker/internal/entity"
	"example.com/tracker/internal/handler/api"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Start timer for task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param        id   path      int  true  "User ID"
// @Param data body api.TaskStartRequest true "Request body"
//
// @Success 200 {object} api.DefaultResponse "Task timer started succesfully"
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /tasks/start/{id} [post]
func (h *Handler) startTimer(ctx *gin.Context) {
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

	var req api.TaskStartRequest
	err = ctx.ShouldBindJSON(&req)

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	task := &entity.Task{
		UserID:      reqID.Value,
		Name:        req.TaskName,
		Description: req.TaskDescription,
		EndAt:       req.TaskEnds,
	}
	logrus.Debugf("Starting task for userID:%d", reqID.Value)
	err = h.Service.StartTimer(task)

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.Header("Locations", fmt.Sprintf("/tasks/%d", task.ID))
	ctx.JSON(http.StatusOK, &api.DefaultResponse{
		Code:    http.StatusOK,
		Message: "task timer started succesfully",
	})
}

// @Summary End timer for task
// @Tags Tasks
// @Produce json
// @Param id path int true "Task ID"
//
// @Success 200 {object} api.DefaultResponse "Task timer succesfully ended"
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /tasks/end/{id} [post]
func (h *Handler) endTimer(ctx *gin.Context) {
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

	logrus.Debugf("Ending task ID:%d", reqID.Value)
	err = h.Service.EndTimer(reqID.Value)

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
		Message: "task timer succesfully ended",
	})
}

// @Summary Get labor costs by userID
// @Tags Tasks
// @Accept json
// @Produce json
// @Param        id   path      int  true  "User ID"
// @Param data query api.WorkLoadsRequest true "Request body"
//
// @Success 200 {object} api.WorkLoadsResponse
// @Failure 400 {object} api.ErrorResponse
// @Failure 500 {object} api.ErrorResponse
// @Router /tasks/workload/{id} [get]
func (h *Handler) getWorkLoads(ctx *gin.Context) {
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

	var req api.WorkLoadsRequest
	err = ctx.ShouldBindQuery(&req)

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	workload, err := h.Service.GetWorkLoads(reqID.Value, req.StartDate, req.EndDate)

	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, &api.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.WorkLoadsResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Body:    workload,
	})
}
