package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/thalysonalexr/movie-poster/usecase"
)

func createError(err error, status int) []byte {
	e, _ := json.Marshal(ErrorPresenter{
		Type:    "error",
		Status:  status,
		Message: err.Error(),
	})
	return e
}

// GetMovies handler to get all movies
func GetMovies(s usecase.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		gender := r.URL.Query().Get("gender")
		movies, err := s.SearchByGender(gender)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			e := createError(err, http.StatusInternalServerError)
			w.Write(e)
			return
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
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			e := createError(err, http.StatusInternalServerError)
			w.Write(e)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
	}
}
