package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/thalysonalexr/movie-poster/usecase"
)

// GetMovies handler to get all movies
func GetMovies(s usecase.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		gender := r.URL.Query().Get("gender")
		movies, err := s.SearchByGender(gender)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(`{"error": "error to load movies"}`))
		}
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "error to encode data"}`))
		}
		var toPresenter []MoviePresenter
		for _, movie := range movies {
			toPresenter = append(toPresenter, MoviePresenter{
				Title:       movie.Title,
				ReleaseDate: time.Unix(int64(movie.ReleaseDate), 1000),
				Genres:      movie.Genres,
			})
		}
		encoded, err := json.Marshal(toPresenter)
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
	}
}
