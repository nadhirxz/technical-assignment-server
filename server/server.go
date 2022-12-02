package server

import (
	"net/http"
	"os"
	"strconv"

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

	router.GET("/users/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

		user, err := model.GetUser(id)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	router.POST("/users/new", func(c *gin.Context) {
		var user model.User

		err := c.BindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
			})
			return
		}

		err = model.CreateUser(user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created",
			"user":    user,
		})
	})

	router.Run("localhost:" + os.Getenv("PORT"))
}
