package main

import (
	"net/http"
	"sliceMC/business"
	"time"

	"github.com/gin-gonic/gin"
)

var hotelManager *business.HotelManager
var bookingService *business.BookingService

func ExpireHeldBookings() {
	for {
		bookingService.CheckExpiredBookings()
		time.Sleep(10 * time.Minute)
	}
}

func main() {
	hotelManager = business.NewHotelManager()
	bookingService = business.NewBookingService(hotelManager)
	router := gin.Default()
	api := router.Group("")
	api.POST("api/rooms/create", createRoom)
	api.GET("/api/rooms/status", getAllRoomsByStatus)

	api.POST("api/booking/create", createBooking)
	api.POST("api/booking/cancel", cancelBooking)
	api.POST("api/booking/confirm", confirmBooking)

	router.Run("localhost:8080")

	go ExpireHeldBookings()
}

type CreateRoomRequest struct {
	RoomType int `json:"roomType"`
}

func createRoom(c *gin.Context) {
	var req CreateRoomRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Payload",
			"error":   err.Error(),
		})
		return
	}
	roomId := hotelManager.CreateRoom(req.RoomType)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created room",
		"roomId":  roomId,
	})
}

type CreateBookingRequest struct {
	RoomId int `json:"roomId"`
	UserId int `json:"userId"`
}

func createBooking(c *gin.Context) {
	var req CreateBookingRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Payload",
			"error":   err.Error(),
		})
		return
	}
	bookingId, err := bookingService.CreateBooking(uint64(req.UserId), uint64(req.RoomId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create booking",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "Successfully created booking",
		"bookingId": bookingId,
	})
}

type CancelBookingRequest struct {
	BookingId int `json:"bookingId"`
}

func cancelBooking(c *gin.Context) {
	var req CancelBookingRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Payload",
			"error":   err.Error(),
		})
		return
	}
	err = bookingService.CancelBooking(uint64(req.BookingId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to cancel booking",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully cancelled booking",
	})
}

type ConfirmBookingRequest struct {
	BookingId int `json:"bookingId"`
}

func confirmBooking(c *gin.Context) {
	var req CancelBookingRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Payload",
			"error":   err.Error(),
		})
		return
	}
	err = bookingService.ConfirmBooking(uint64(req.BookingId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to confirm booking",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully confirm booking",
	})
}

type GetRoomsRequest struct {
	RoomStatus int `form:"status"`
}

func getAllRoomsByStatus(c *gin.Context) {
	var req GetRoomsRequest
	err := c.BindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid query param",
			"error":   err.Error(),
		})
		return
	}
	rooms := hotelManager.GetAllRoomsByStatus(req.RoomStatus)
	c.JSON(http.StatusOK, rooms)

}
