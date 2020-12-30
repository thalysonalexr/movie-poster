package api

import "time"

// MoviePresenter struct
type MoviePresenter struct {
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Genres      []string  `json:"genres"`
}
