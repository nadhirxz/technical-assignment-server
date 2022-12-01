package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	router.Run("localhost:" + os.Getenv("PORT"))
}
