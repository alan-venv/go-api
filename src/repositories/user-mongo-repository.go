package repositories

import (
	"context"
	"errors"
	"example/go-api/src/models"
	"time"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	Database *mongo.Database
}

// ! ========================================
func (self UserMongoRepository) ReadAll() ([]models.User, error) {
	db := self.Database
	collection := db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []models.User

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	err = cursor.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// ! ========================================
func (self UserMongoRepository) Read(id string) (models.User, error) {
	db := self.Database
	collection := db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// ! ========================================
func (self UserMongoRepository) Create(user models.User) error {
	db := self.Database
	collection := db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.ID = uuid.New().String()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// ! ========================================

// ! ========================================
func (self UserMongoRepository) Delete(id string) error {
	db := self.Database
	collection := db.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, _ := collection.DeleteOne(ctx, bson.M{"_id": id})

	if result.DeletedCount == 0 {
		return errors.New("Cannot find or delete user")
	}

	return nil
}
