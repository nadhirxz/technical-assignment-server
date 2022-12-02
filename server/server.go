package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/nadhirxz/technical-assignment-server/model"
)

func Init() {
	router := gin.Default()
	model.Init()

	router.GET("/users", func(c *gin.Context) {
		users, err := model.GetUsers()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, users)
	})

	router.Run("localhost:" + os.Getenv("PORT"))
}
