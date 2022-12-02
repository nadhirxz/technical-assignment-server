package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/nadhirxz/technical-assignment-server/model"
)

func Init() {
	router := gin.Default()
	model.Init()

	router.Run("localhost:" + os.Getenv("PORT"))
}
