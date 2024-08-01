package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var messages MessageList
var groups map[string]*Group

func main() {
	messages = NewMessageList()
	groups = make(map[string]*Group)

	router := gin.Default()
	router.GET("/receiveMessage", receiveMessage)
	router.POST("/sendMessage", sendMessage)
	router.POST("/createGroup", createGroup)
	router.POST("/addUserToGroup", addUserToGroup)
	router.POST("/sendMessageToGroup", sendMessageToGroup)
	router.GET("/getMessagesFromGroup", getMessagesFromGroup)

	router.Run("localhost:8080")

}

type sendRequest struct {
	Message  string `json:'message'`
	Username string
}

func sendMessage(c *gin.Context) {
	var req sendRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	log.Println(req)
	messages.AddMessage(req.Message)
	c.IndentedJSON(http.StatusAccepted, "message received successfully")
}

type receiveRequest struct {
	username string `form:"username"`
}

func receiveMessage(c *gin.Context) {
	var req receiveRequest
	if err := c.BindQuery(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	msgs := messages.GetAllMessages()
	c.IndentedJSON(http.StatusAccepted, msgs)

}

type CreateGroupReq struct {
	GroupName string `json:"groupName"`
}

func createGroup(c *gin.Context) {
	var req CreateGroupReq
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	g := NewGroup(req.GroupName)
	groups[req.GroupName] = g
	log.Println(groups)
	c.IndentedJSON(http.StatusAccepted, "group created successfully")
}

type AddUserToGroupReq struct {
	GroupName string `json:"groupName"`
	UserName  string `json:"userName"`
}

func addUserToGroup(c *gin.Context) {
	var req AddUserToGroupReq
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if _, ok := groups[req.GroupName]; !ok {
		c.IndentedJSON(http.StatusBadRequest, "this group doesn't exist")
		return
	}
	groups[req.GroupName].AddUser(req.UserName)
	c.IndentedJSON(http.StatusAccepted, "user added to group successfully")
}

type AddMessageToGroupReq struct {
	GroupName string `json:"groupName"`
	UserName  string `json:"userName"`
	Message   string `json:"message"`
}

func sendMessageToGroup(c *gin.Context) {
	var req AddMessageToGroupReq
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if _, ok := groups[req.GroupName]; !ok {
		c.IndentedJSON(http.StatusBadRequest, "this group doesn't exist")
		return
	}
	if err := groups[req.GroupName].AddMessage(req.UserName, req.Message); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusAccepted, "message added to group successfully")
}

type GetMessagesFromGroupRequest struct {
	GroupName string `form:"groupName"`
	UserName  string `form:"userName"`
}

func getMessagesFromGroup(c *gin.Context) {
	var req GetMessagesFromGroupRequest
	if err := c.BindQuery(&req); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	log.Println(req.GroupName, groups)
	if _, ok := groups[req.GroupName]; !ok {
		c.IndentedJSON(http.StatusBadRequest, "this group doesn't exist")
		return
	}
	if msgs, err := groups[req.GroupName].GetAllMessages(req.UserName); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	} else {
		log.Println(msgs)
		c.JSON(http.StatusAccepted, msgs)
	}

}
