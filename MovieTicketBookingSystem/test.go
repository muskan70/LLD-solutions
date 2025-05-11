package main

import (
	"errors"
	"fmt"
	"ticketBooking/constants"
	"ticketBooking/services"
)

func GetFirstNAvailableSeats(seats map[string]int, n int) ([]string, error) {
	var seatsToBook []string
	for id, status := range seats {
		if status == constants.SEAT_STATUS_AVAILABLE {
			seatsToBook = append(seatsToBook, id)
		}
		if len(seatsToBook) == n {
			return seatsToBook, nil
		}
	}
	return nil, errors.New("seats not available")
}

func test() {
	theatreId := theatreManager.CreateTheatre("Pebble Downtown", "Faridabad", "sector-12, 121006")
	screenId := theatreManager.CreateScreen("Audi1", theatreId, 5, 7)
	movieId := movieManager.CreateMovie("Raid2", "Story of Income tax raid by Amay Patnayak", 120, []int{constants.LANGUAGE_HINDI, constants.LANGUAGE_ENGLISH}, constants.GENRE_ACTION, "02-05-2025")
	fmt.Println("theatreId created:", theatreId, "\nscreenId created:", screenId, "\nmovieId created:", movieId)

	showId1 := showManager.CreateShow(movieId, screenId, "13:00", "11-05-2025", map[int]float64{
		constants.SEAT_CATEGORY_SILVER:   120,
		constants.SEAT_CATEGORY_PLATINUM: 150,
	}, constants.LANGUAGE_HINDI)

	showId2 := showManager.CreateShow(movieId, screenId, "16:00", "11-05-2025", map[int]float64{
		constants.SEAT_CATEGORY_SILVER:   120,
		constants.SEAT_CATEGORY_PLATINUM: 150,
	}, constants.LANGUAGE_ENGLISH)

	showId3 := showManager.CreateShow(movieId, screenId, "20:00", "11-05-2025", map[int]float64{
		constants.SEAT_CATEGORY_SILVER:   120,
		constants.SEAT_CATEGORY_PLATINUM: 150,
	}, constants.LANGUAGE_HINDI)
	fmt.Println("showIds created:", showId1, showId2, showId3)

	userId1 := userManager.CreateUser("muskan", "muskanmanglaxxx@gmail.com", 999999999)
	fmt.Println("userId created:", userId1)

	lang := constants.LANGUAGE_HINDI
	searchResults := searchService.SearchMovie(&services.SearchParams{
		City:     "Faridabad",
		Date:     "11-05-2025",
		Language: &lang,
	})

	fmt.Println(searchResults)

	seats := showManager.GetShowById(searchResults[movieId][0]).GetSeatsLayoutWithStatus()
	fmt.Println(seats)

	seatsToBook, err := GetFirstNAvailableSeats(seats, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		bookingId, err := bookingService.BookTickets(userId1, searchResults[movieId][0], seatsToBook)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("bookingId created:", bookingId)
			bookingService.CancelTickets(bookingId)
		}
	}
}
