package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mendezdev/golang_mongo-example/domain/users"
	"github.com/mendezdev/golang_mongo-example/services"
	"github.com/mendezdev/golang_mongo-example/utils/api_errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := api_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userID, userIdErr := getUserID(c)
	if userIdErr != nil {
		c.JSON(userIdErr.Status(), userIdErr)
		return
	}

	user, getErr := services.UsersService.GetUser(*userID)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	err := api_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
	c.JSON(err.Status(), err)
}

func Delete(c *gin.Context) {
	userID, userIdErr := getUserID(c)
	if userIdErr != nil {
		c.JSON(userIdErr.Status(), userIdErr)
		return
	}

	deleteErr := services.UsersService.DeleteUser(*userID)
	if deleteErr != nil {
		c.JSON(deleteErr.Status(), deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func getUserID(c *gin.Context) (*string, api_errors.RestErr) {
	userID := c.Param("user_id")
	if userID == "" {
		return nil, api_errors.NewBadRequestError("user id should be a number")
	}

	return &userID, nil
}
