package main

import (
	"net/http"

	"github.com/thalysonalexr/movie-poster/api"
	repository "github.com/thalysonalexr/movie-poster/infra/repo"
	service "github.com/thalysonalexr/movie-poster/usecase"
)

func main() {
	repo := repository.MoviesRepositoryImpl{}
	service := service.CreateNewService(&repo)

	http.HandleFunc("/movies", api.GetMovies(service))
	http.ListenAndServe(":8080", nil)
}
