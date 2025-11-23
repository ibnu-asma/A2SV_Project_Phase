package repositories

import (
	"context"
	"errors"
	"task_manager/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	GetAll() ([]domain.Task, error)
	GetByID(id string) (domain.Task, error)
	Create(task domain.Task) (domain.Task, error)
	Update(id string, task domain.Task) (domain.Task, error)
	Delete(id string) error
}

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(client *mongo.Client, dbName, collectionName string) TaskRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &taskRepository{collection: collection}
}

func (r *taskRepository) GetAll() ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []domain.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	if tasks == nil {
		tasks = []domain.Task{}
	}
	return tasks, nil
}

func (r *taskRepository) GetByID(id string) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, errors.New("invalid task ID")
	}

	var task domain.Task
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return domain.Task{}, errors.New("task not found")
	}
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) Create(task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	task.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, task)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) Update(id string, updatedTask domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, errors.New("invalid task ID")
	}

	updatedTask.ID = objectID
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return domain.Task{}, err
	}
	if result.MatchedCount == 0 {
		return domain.Task{}, errors.New("task not found")
	}

	return updatedTask, nil
}

func (r *taskRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
