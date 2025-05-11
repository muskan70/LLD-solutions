package main

import (
	"ticketBooking/manager"
	"ticketBooking/services"
)

var movieManager *manager.MovieManager
var theatreManager *manager.TheatreManager
var showManager *manager.ShowManager
var userManager *manager.UserManager

var bookingService *services.BookingService
var searchService *services.SearchCatalogue

func InitializeManagers() {
	movieManager = manager.NewMovieManager()
	theatreManager = manager.NewTheatreManager()
	showManager = manager.NewShowManager()
	userManager = manager.NewUserManager()
}

func InitializeServices() {
	bookingService = services.NewBookingService(showManager)
	searchService = services.NewSearchService(movieManager, theatreManager, showManager)
}
