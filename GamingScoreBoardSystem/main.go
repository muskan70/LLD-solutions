package main

import (
	"database/sql"
	"intuitMc/domain/user"
	"intuitMc/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	router := gin.Default()

	router.GET("/topScores", getTopScores)
	router.POST("/register/user", registerUser)
	router.Run("localhost:8080")
}

func getTopScores(c *gin.Context) {
	users, err := user.GetTopKUserScores(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "users": users})
}

func registerUser(c *gin.Context) {
	var usr *requests.RegisterUserRequest
	err := c.BindJSON(&usr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}
	if err := user.RegisterUser(c, usr); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true})
}
