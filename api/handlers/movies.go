package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/thalysonalexr/movie-poster/api/presenters"
	"github.com/thalysonalexr/movie-poster/usecase"
)

func createError(err error, status int) []byte {
	e, _ := json.Marshal(presenters.ErrorPresenter{
		Type:    "error",
		Status:  status,
		Message: err.Error(),
	})
	return e
}

func listMovies(s usecase.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		gender := r.URL.Query().Get("gender")
		movies, err := s.SearchByGender(gender)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			e := createError(err, http.StatusInternalServerError)
			w.Write(e)
			return
		}
		var toPresenter []presenters.MoviePresenter
		for _, movie := range movies {
			toPresenter = append(toPresenter, presenters.MoviePresenter{
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
	})
}

func downloadPosters(s usecase.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gender := r.URL.Query().Get("gender")
		success, err := s.DownloadPosters(gender)
		if err != nil || !success {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

// MakeMovieHandlers make url handlers
func MakeMovieHandlers(r *mux.Router, n negroni.Negroni, service usecase.Service) {
	r.Handle("/movies", n.With(
		negroni.Wrap(listMovies(service)),
	)).Methods("GET", "OPTIONS").Name("listMovies")

	r.Handle("/movies/download-posters", n.With(
		negroni.Wrap(downloadPosters(service)),
	)).Methods("GET", "OPTIONS").Name("downloadPosters")
}
