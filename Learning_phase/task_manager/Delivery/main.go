package main

import (
	"context"
	"log"
	"task_manager/Delivery/controllers"
	"task_manager/Delivery/routers"
	"task_manager/Infrastructure"
	"task_manager/Repositories"
	"task_manager/Usecases"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(ctx)

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Connected to MongoDB successfully!")

	taskRepo := repositories.NewTaskRepository(client, "taskdb", "tasks")
	userRepo := repositories.NewUserRepository(client, "taskdb", "users")

	passwordService := infrastructure.NewPasswordService()
	jwtService := infrastructure.NewJWTService()

	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo, passwordService, jwtService)

	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	authMiddleware := infrastructure.NewAuthMiddleware(jwtService)

	r := routers.SetupRouter(taskController, userController, authMiddleware)
	r.Run(":8080")
}
