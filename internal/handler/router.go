package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")

	users.POST("/create", h.createUser)
	users.GET("/", h.getUsers)
	users.DELETE("/delete/:id", h.deleteUser)
	users.PATCH("/update/:id", h.updateUser)
	return router
}
