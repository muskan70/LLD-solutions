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
	router.POST("/user/login", loginUser)
	router.GET("/receiveMessage", receiveMessage)
	router.POST("/sendMessage", sendDirectMessage)
	router.GET("/chatHistory", getChatHistory)
	router.POST("/group/create", createGroup)
	router.POST("/group/addUser", addUserToGroup)
	router.POST("/group/removeUser", removeUserFromGroup)
	router.POST("/group/sendMessage", sendMessageToGroup)
	router.GET("/group/getAllMessages", getMessagesFromGroup)

	router.Run("localhost:8080")

}

type loginUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone int    `json:"phone"`
}

func loginUser(c *gin.Context) {
	var req loginUserRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	usrId := chatSystem.AddNewUser(req)
	c.IndentedJSON(http.StatusAccepted, gin.H{
		"msg":    "user login successfully",
		"userId": usrId,
	})
}

type sendDirectMessageRequest struct {
	Content    string `json:"message"`
	SenderId   int    `json:"senderId"`
	ReceiverId int    `json:"receiverId"`
}

func sendDirectMessage(c *gin.Context) {
	var req sendDirectMessageRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	chatSystem.HandleDirectMessage(req)
	c.IndentedJSON(http.StatusAccepted, "message sent successfully")
}

type receiveRequest struct {
	UserId int `form:"userId"`
}

func receiveMessage(c *gin.Context) {
	var req receiveRequest
	if err := c.BindQuery(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(req)
	directMsgs, grpMsgs, err := chatSystem.ReceiveMessages(req.UserId)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{
		"direct Messages": directMsgs,
		"group Messages":  grpMsgs,
	})
}

func getChatHistory(c *gin.Context) {
	var req receiveRequest
	if err := c.BindQuery(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(req)
	directMsgs, grpMsgs, err := chatSystem.ChatHistoryOfUser(req.UserId)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{
		"direct Messages": directMsgs,
		"group Messages":  grpMsgs,
	})

}

type CreateGroupReq struct {
	GroupName string `json:"groupName"`
	AdminId   int    `json:"adminId"`
}

func createGroup(c *gin.Context) {
	var req CreateGroupReq
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	groupId, err := chatSystem.CreateGroup(req)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
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
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := chatSystem.AddUserToGroup(req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
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
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := chatSystem.RemoveUserFromGroup(req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
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
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := chatSystem.HandleGroupMessage(req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
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
		c.IndentedJSON(http.StatusBadRequest, err.Error())
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
