package main

import (
	"intuitMc/cron"
	"intuitMc/domain/score"
	"intuitMc/domain/user"
	"intuitMc/requests"
	"intuitMc/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		c.JSON(http.StatusOK, response.Fail("failed to get top scores", err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithTopScores("successfully fetched top scores", users))
}

func registerUser(c *gin.Context) {
	var usr *requests.RegisterUserRequest
	err := c.BindJSON(&usr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("bad request", err))
		return
	}
	userId, err := user.RegisterUser(c, usr)
	if err != nil {
		c.JSON(http.StatusOK, response.Fail("failed to register User", err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithUserId("successfully registered user", userId))
}

func pushScoresToFile(c *gin.Context) {
	var scr *requests.PushScoreRequest
	err := c.BindJSON(&scr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Fail("bad requerst", err))
		return
	}
	if err := score.PushScoreToFile(c, scr); err != nil {
		c.JSON(http.StatusOK, response.Fail("failed to push scores to File", err))
		return
	}
	c.JSON(http.StatusOK, response.Success("Scores pushed successfully"))
}
