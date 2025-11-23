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

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	CountUsers() (int64, error)
	PromoteToAdmin(username string) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, dbName, collectionName string) UserRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &userRepository{collection: collection}
}

func (r *userRepository) Create(user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetByUsername(username string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user domain.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return domain.User{}, errors.New("user not found")
	}
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *userRepository) CountUsers() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.M{})
	return count, err
}

func (r *userRepository) PromoteToAdmin(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.collection.UpdateOne(
		ctx,
		bson.M{"username": username},
		bson.M{"$set": bson.M{"role": "admin"}},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}
