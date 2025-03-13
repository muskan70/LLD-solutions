package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var chats map[int]*Chat
var groups map[int]*Group

func main() {
	chats = make(map[int]*Chat)
	groups = make(map[int]*Group)

	router := gin.Default()
	//router.GET("/receiveMessage", receiveMessage)
	//router.POST("/sendMessage", sendMessage)
	router.POST("/createGroup", createGroup)
	router.POST("/addUserToGroup", addUserToGroup)
	router.POST("/removeUserfromGroup", removeUserFromGroup)
	router.POST("/sendMessageToGroup", sendMessageToGroup)
	router.GET("/getMessagesFromGroup", getMessagesFromGroup)

	router.Run("localhost:8080")

}

// type sendRequest struct {
// 	Msg      string `json:'message'`
// 	Username string
// }

// func sendMessage(c *gin.Context) {
// 	var req sendRequest
// 	if err := c.BindJSON(&req); err != nil {
// 		log.Println(err)
// 		c.IndentedJSON(http.StatusBadRequest, err)
// 		return
// 	}
// 	log.Println(req)
// 	messages.AddMessage(req.Message)
// 	c.IndentedJSON(http.StatusAccepted, "message received successfully")
// }

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
	g := NewGroup(req.GroupName, req.AdminId)
	groups[g.GroupId] = g
	log.Println(groups)
	c.IndentedJSON(http.StatusAccepted, gin.H{
		"GroupId": g.GroupId,
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
	if _, ok := groups[req.GroupId]; !ok {
		c.IndentedJSON(http.StatusBadRequest, "this group doesn't exist")
		return
	}
	groups[req.GroupId].AddParticipant(req.UserId)
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
	if _, ok := groups[req.GroupId]; !ok {
		c.IndentedJSON(http.StatusBadRequest, "this group doesn't exist")
		return
	}
	groups[req.GroupId].RemoveParticipant(req.UserId)
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
	if _, ok := groups[req.GroupId]; !ok {
		c.IndentedJSON(http.StatusBadRequest, "this group doesn't exist")
		return
	}
	if err := groups[req.GroupId].AddMessage(req.UserId, req.Content); err != nil {
		log.Println(err)
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

	if _, ok := groups[req.GroupId]; !ok {
		c.IndentedJSON(http.StatusBadRequest, "this group doesn't exist")
		return
	}
	if msgs, err := groups[req.GroupId].GetChatHistory(req.UserId); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	} else {
		log.Println(msgs)
		c.JSON(http.StatusAccepted, msgs)
	}
}
