package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	router "github.com/gorilla/mux"
	"github.com/thalysonalexr/movie-poster/api/handlers"
	"github.com/thalysonalexr/movie-poster/api/middlewares"
	repository "github.com/thalysonalexr/movie-poster/infra/repo"
	service "github.com/thalysonalexr/movie-poster/usecase"
)

func main() {
	repo := repository.MoviesRepositoryImpl{}
	service := service.CreateNewService(&repo)

	r := router.NewRouter()
	n := negroni.New(
		negroni.HandlerFunc(middlewares.Cors),
		negroni.NewLogger(),
	)

	handlers.MakeMovieHandlers(r, *n, service)
	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
