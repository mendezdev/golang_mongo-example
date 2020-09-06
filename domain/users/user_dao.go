package users

import (
	"context"
	"fmt"

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
		fmt.Println("error when trying to insert User", insertErr)
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
		fmt.Println("error when trying to parse ID to get user in db", userIDErr)
		return api_errors.NewBadRequestError("invalid id to get user")
	}

	filter := bson.D{{"_id", userID}}

	userGetErr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if userGetErr != nil {
		if userGetErr == mongo.ErrNoDocuments {
			return api_errors.NewNotFoundError("user not found with given id")
		}
		fmt.Println("error trying to get user with given filter", userGetErr)
		return api_errors.NewInternalServerError("database error", userGetErr)
	}

	return nil
}

func (user *User) GetByEmail() api_errors.RestErr {
	collection := getUserCollection()

	filter := bson.D{{"email", user.Email}}

	userGetErr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if userGetErr != nil {
		if userGetErr == mongo.ErrNoDocuments {
			return nil
		}
		fmt.Println("error trying to get user by email", userGetErr)
		return api_errors.NewInternalServerError("database error", userGetErr)
	}

	return nil
}

func findOneByFilter(fieldName string, user *User) (*User, api_errors.RestErr) {
	collection := getUserCollection()

	filter := bson.D{{fieldName, user.Email}}

	userGetErr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if userGetErr != nil {
		if userGetErr == mongo.ErrNoDocuments {
			return nil, api_errors.NewNotFoundError(fmt.Sprintf("user not found with given filter: %s", fieldName))
		}
		fmt.Println("error trying to get user with given filter", userGetErr)
		return nil, api_errors.NewInternalServerError("database error", userGetErr)
	}

	return user, nil
}

func getUserCollection() *mongo.Collection {
	return mongodb.MongoClient.Database(DB_NAME).Collection(USERS_COLLECTION_NAME)
}
