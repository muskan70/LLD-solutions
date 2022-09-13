package main

import (
	"database/sql"
	"intuitMc/cron"
	"intuitMc/domain/score"
	"intuitMc/domain/user"
	"intuitMc/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	router := gin.Default()
	cron.ScoreProcessing()

	router.GET("/topScores", getTopScores)
	router.POST("/register/user", registerUser)
	router.POST("/score/push", pushScoresToFile)
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

func pushScoresToFile(c *gin.Context) {
	var scr *requests.PushScoreRequest
	err := c.BindJSON(&scr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}
	if err := score.PushScoreToFile(c, scr); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true})
}
