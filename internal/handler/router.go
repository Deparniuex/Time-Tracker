package handler

import (
	_ "example.com/tracker/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	users := router.Group("/users")
	tasks := router.Group("/tasks")

	users.POST("/create", h.createUser)
	users.GET("/", h.getUsers)
	users.DELETE("/delete/:id", h.deleteUser)
	users.PUT("/update/:id", h.updateUser)

	tasks.POST("/start/:id", h.startTimer)
	tasks.POST("/end/:id", h.endTimer)
	tasks.GET("/workload/:id", h.getWorkLoads)
	return router
}
