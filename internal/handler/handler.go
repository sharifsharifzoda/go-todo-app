package handler

import (
	"github.com/gin-gonic/gin"
	"todo_sql_database/internal/service"
	"todo_sql_database/logging"
)

type Handler struct {
	Auth   service.Authorization
	Todo   service.TodoTask
	Logger *logging.Logger
}

func NewHandler(auth service.Authorization, todo service.TodoTask, log *logging.Logger) *Handler {
	return &Handler{
		Auth:   auth,
		Todo:   todo,
		Logger: log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	mainRoute := router.Group("/main")

	auth := mainRoute.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := mainRoute.Group("/api", h.tokenAuthMiddleware)
	{
		api.POST("/task", h.createTask)
		api.GET("/task", h.getTasks)
		api.GET("/task/:id", h.getTaskById)
		api.PUT("/task/:id", h.updateTask)
		api.DELETE("/task/:id", h.deleteTask)
	}

	return router
}
