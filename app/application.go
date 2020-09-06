package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mendezdev/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

//StartApplication set all the url mappings and start the server
func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":8080")
}
