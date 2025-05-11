package services

import "time"

type SeachCatalogue struct {
	MovieManager   *MovieManager
	TheatreManager *TheatreManager
	ShowManager    *ShowManager
}

type SearchParams struct {
	Genre    int
	Language int
	City     string
	Title    string
	Date     time.Time
}

func (s *SeachCatalogue) SearchMovie(scrh *SearchParams) {

}
