package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(client *mongo.Client, dbName, collectionName string) *UserService {
	collection := client.Database(dbName).Collection(collectionName)
	return &UserService{collection: collection}
}

func (s *UserService) Register(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existing, _ := s.GetUserByUsername(user.Username)
	if existing.Username != "" {
		return models.User{}, errors.New("username already exists")
	}

	count, err := s.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return models.User{}, err
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)

	_, err = s.collection.InsertOne(ctx, user)
	if err != nil {
		return models.User{}, err
	}

	user.Password = ""
	return user, nil
}

func (s *UserService) Login(username, password string) (models.User, error) {
	user, err := s.GetUserByUsername(username)
	if err != nil {
		return models.User{}, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("invalid credentials")
	}

	user.Password = ""
	return user, nil
}

func (s *UserService) GetUserByUsername(username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := s.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return models.User{}, errors.New("user not found")
	}
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *UserService) PromoteUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := s.collection.UpdateOne(
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
