package routers

import (
	"task_manager/Delivery/controllers"
	"task_manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController, authMiddleware *infrastructure.AuthMiddleware) *gin.Engine {
	r := gin.Default()

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	r.GET("/tasks", authMiddleware.AuthRequired(), taskController.GetTasks)
	r.GET("/tasks/:id", authMiddleware.AuthRequired(), taskController.GetTask)

	admin := r.Group("/")
	admin.Use(authMiddleware.AuthRequired(), authMiddleware.AdminOnly())
	{
		admin.POST("/tasks", taskController.CreateTask)
		admin.PUT("/tasks/:id", taskController.UpdateTask)
		admin.DELETE("/tasks/:id", taskController.DeleteTask)
		admin.PUT("/promote/:username", userController.Promote)
	}

	return r
}
