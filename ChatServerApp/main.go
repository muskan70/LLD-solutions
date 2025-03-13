package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var chatSystem *ChatSystem

func main() {
	chatSystem = NewChatSystem()

	router := gin.Default()
	//router.GET("/receiveMessage", receiveMessage)
	router.POST("/sendMessage", sendMessage)
	router.POST("/createGroup", createGroup)
	router.POST("/addUserToGroup", addUserToGroup)
	router.POST("/removeUserfromGroup", removeUserFromGroup)
	router.POST("/sendMessageToGroup", sendMessageToGroup)
	router.GET("/getMessagesFromGroup", getMessagesFromGroup)

	router.Run("localhost:8080")

}

type sendRequest struct {
	Content    string `json:"message"`
	SenderId   int    `json:"senderId"`
	ReceiverId int    `json:"receiverId"`
}

func sendMessage(c *gin.Context) {
	var req sendRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	chatSystem.HandleDirectMessage(req)
	c.IndentedJSON(http.StatusAccepted, "message received successfully")
}

// type receiveRequest struct {
// 	username string `form:"username"`
// }

// func receiveMessage(c *gin.Context) {
// 	var req receiveRequest
// 	if err := c.BindQuery(&req); err != nil {
// 		log.Println(err)
// 		c.IndentedJSON(http.StatusBadRequest, err)
// 		return
// 	}
// 	msgs := messages.GetAllMessages()
// 	c.IndentedJSON(http.StatusAccepted, msgs)

// }

type CreateGroupReq struct {
	GroupName string `json:"groupName"`
	AdminId   int    `json:"adminId"`
}

func createGroup(c *gin.Context) {
	var req CreateGroupReq
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	groupId := chatSystem.CreateGroup(req)
	c.IndentedJSON(http.StatusAccepted, gin.H{
		"GroupId": groupId,
		"message": "group created successfully",
	})
}

type AddUserToGroupReq struct {
	GroupId int `json:"groupId"`
	UserId  int `json:"userId"`
}

func addUserToGroup(c *gin.Context) {
	var req AddUserToGroupReq
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := chatSystem.AddUserToGroup(req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusAccepted, "user added to group successfully")
}

type RemoveUserFromGroupReq struct {
	GroupId int `json:"groupId"`
	UserId  int `json:"userId"`
}

func removeUserFromGroup(c *gin.Context) {
	var req RemoveUserFromGroupReq
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := chatSystem.RemoveUserFromGroup(req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusAccepted, "user removed from group successfully")
}

type AddMessageToGroupReq struct {
	GroupId int    `json:"groupId"`
	UserId  int    `json:"userId"`
	Content string `json:"message"`
}

func sendMessageToGroup(c *gin.Context) {
	var req AddMessageToGroupReq
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := chatSystem.HandleGroupMessage(req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusAccepted, "message sent to group successfully")
}

type GetMessagesFromGroupRequest struct {
	GroupId int `form:"groupId"`
	UserId  int `form:"userId"`
}

func getMessagesFromGroup(c *gin.Context) {
	var req GetMessagesFromGroupRequest
	if err := c.BindQuery(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if msgs, err := chatSystem.GetGroupChatHistory(req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	} else {
		log.Println(msgs)
		c.JSON(http.StatusAccepted, msgs)
	}
}
