package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")

	user.GET("/", h.getUsers)
	return router
}
