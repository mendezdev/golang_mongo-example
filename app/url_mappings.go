package app

import "github.com/mendezdev/golang_mongo-example/controllers/ping"

func mapUrls() {
	router.GET("/ping", ping.Ping)
}
