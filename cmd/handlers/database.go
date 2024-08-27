// This files contains helper function used with database.
// These functions help us to reuse the code and also helps in avoiding code redendency
package handlers

import (
	"context"
	"errors"
	"time"

	config "github.com/ShikharY10/goauth/cmd/configs"
	"github.com/ShikharY10/goauth/cmd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataBase struct {
	MongoDB *config.MongoDB
}

// Initializes Database handler
func CreateDatabaseHandler(mongodb *config.MongoDB) *DataBase {
	return &DataBase{
		MongoDB: mongodb,
	}
}

// Create a new document in mongodb of {models.User} type.
func (db *DataBase) CreateNewUser(user models.User) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()
	_, err := db.MongoDB.User.InsertOne(ctx, user)
	return err
}

// Return {models.User} based on filter and search options
func (db *DataBase) GetUserData(filter interface{}, findOptions *options.FindOptions) (*models.User, error) {
	cursor, err := db.MongoDB.User.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			continue
		} else {
			users = append(users, user)
		}
	}
	if len(users) > 0 {
		return &users[0], nil
	}
	return nil, errors.New("no document found")
}

// return an array of models.User based on filter and search result
func (db *DataBase) GetMultipleUsers(filter interface{}, findOptions options.FindOptions, includePassword bool) ([]models.User, error) {

	cursor, err := db.MongoDB.User.Find(context.TODO(), filter, &findOptions)
	if err != nil {
		return nil, err
	} else {
		var users []models.User

		for cursor.Next(context.TODO()) {
			var user models.User
			if err := cursor.Decode(&user); err != nil {
				continue
			} else {
				if !includePassword {
					user.Password = ""
				}
				users = append(users, user)
			}
		}
		return users, nil
	}
}

// Delete a user document from database based on filter
func (db *DataBase) DeleteUser(filter interface{}) (*models.User, error) {
	cursor := db.MongoDB.User.FindOneAndDelete(context.TODO(), filter)
	var user models.User
	err := cursor.Decode(&user)
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

// Changes the role of user
func (db *DataBase) ChangeRole(id string, role string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	} else {
		result, err := db.MongoDB.User.UpdateOne(
			context.TODO(),
			bson.M{"_id": _id},
			bson.M{"$set": bson.M{"role": role}},
		)
		if err != nil {
			return err
		}
		if result.ModifiedCount > int64(0) {
			return nil
		} else {
			return errors.New("no document maatch the filter")
		}
	}
}

// Returns user organisation
func (db *DataBase) GetUserOrganisation(id string) (string, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	} else {
		opts := options.Find().SetProjection(
			bson.D{
				{Key: "_id", Value: 0},
				{Key: "name", Value: 0},
				{Key: "username", Value: 0},
				{Key: "password", Value: 0},
				{Key: "role", Value: 0},
			},
		)
		user, err := db.GetUserData(bson.M{"_id": _id}, opts)
		if err != nil {
			return "", err
		} else {
			return user.Organisation, nil
		}
	}
}
