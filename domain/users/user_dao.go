package users

import (
	"context"
	"fmt"

	"github.com/mendezdev/bookstore_users-api/logger"
	"github.com/mendezdev/golang_mongo-example/db/mongodb"
	"github.com/mendezdev/golang_mongo-example/utils/api_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	//DB_NAME is the db name
	DB_NAME = "mydb"
	//USERS_COLLECTION_NAME is the name for the users collection
	USERS_COLLECTION_NAME = "users"
)

//Save save a new user to the users collections
func (user *User) Save() api_errors.RestErr {
	collection := getUserCollection()
	insertResult, insertErr := collection.InsertOne(context.TODO(), user)

	if insertErr != nil {
		logger.Error("error when trying to insert User", insertErr)
		return api_errors.NewInternalServerError("database error", insertErr)
	}

	fmt.Println("INSERTED ID: ", insertResult)
	return nil
}

//Get get the user by the ID given
func (user *User) Get() api_errors.RestErr {
	collection := getUserCollection()

	userID, userIDErr := primitive.ObjectIDFromHex(user.ID)
	if userIDErr != nil {
		logger.Error("error when trying to parse ID to get user in db", userIDErr)
		return api_errors.NewBadRequestError("invalid id to get user")
	}

	filter := bson.D{{"_id", userID}}

	userGetErr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if userGetErr != nil {
		logger.Error("error trying to get user with given filter", userGetErr)
		return api_errors.NewInternalServerError("database error", userGetErr)
	}

	return nil
}

func getUserCollection() *mongo.Collection {
	return mongodb.MongoClient.Database(DB_NAME).Collection(USERS_COLLECTION_NAME)
}
