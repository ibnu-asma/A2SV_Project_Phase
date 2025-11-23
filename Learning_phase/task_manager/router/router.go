package router

import (
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()

	taskService := data.NewTaskService(client, "taskdb", "tasks")
	userService := data.NewUserService(client, "taskdb", "users")

	taskController := controllers.NewTaskController(taskService)
	userController := controllers.NewUserController(userService)

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	r.GET("/tasks", middleware.AuthMiddleware(), taskController.GetTasks)
	r.GET("/tasks/:id", middleware.AuthMiddleware(), taskController.GetTask)

	admin := r.Group("/")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/tasks", taskController.CreateTask)
		admin.PUT("/tasks/:id", taskController.UpdateTask)
		admin.DELETE("/tasks/:id", taskController.DeleteTask)
		admin.PUT("/promote/:username", userController.Promote)
	}

	return r
}
