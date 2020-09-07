package app

import (
	"github.com/mendezdev/golang_mongo-example/controllers/ping"
	"github.com/mendezdev/golang_mongo-example/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
