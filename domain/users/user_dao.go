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
	isAvailableEmail, emailAvailableErr := user.isAvailableEmail()
	if emailAvailableErr != nil {
		return emailAvailableErr
	}

	//validating email
	if !isAvailableEmail {
		return api_errors.NewBadRequestError("the email provided is not available")
	}
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
	userID, userIDErr := primitive.ObjectIDFromHex(user.ID)
	if userIDErr != nil {
		fmt.Println("error when trying to parse ID to get user in db", userIDErr)
		return api_errors.NewBadRequestError("invalid id to get user")
	}

	findErr := user.findOneByFilter("_id", userID)
	if findErr != nil {
		if findErr == mongo.ErrNoDocuments {
			return api_errors.NewNotFoundError(fmt.Sprintf("user not found with given id: %s", user.ID))
		}
		fmt.Println("error trying to find document", findErr)
		return api_errors.NewInternalServerError("database error", findErr)
	}
	return nil
}

func (user *User) isAvailableEmail() (bool, api_errors.RestErr) {
	findErr := user.findOneByFilter("email", user.Email)
	if findErr != nil {
		if findErr == mongo.ErrNoDocuments {
			return true, api_errors.NewNotFoundError(fmt.Sprintf("user not found with given email: %s", user.Email))
		}
		fmt.Println("error trying to find document", findErr)
		return false, api_errors.NewInternalServerError("database error", findErr)
	}
	return false, nil
}

func (user *User) findOneByFilter(fieldName string, fieldValue interface{}) error {
	collection := getUserCollection()
	filter := bson.D{{fieldName, fieldValue}}

	userGetErr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if userGetErr != nil {
		return userGetErr
	}

	return nil
}

func getUserCollection() *mongo.Collection {
	return mongodb.MongoClient.Database(DB_NAME).Collection(USERS_COLLECTION_NAME)
}
