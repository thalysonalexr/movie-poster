package repo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/thalysonalexr/movie-poster/entity"
	"github.com/thalysonalexr/movie-poster/errors"
)

// Repository interface to repositories
type Repository interface {
	List() ([]entity.Movie, error)
}

// MoviesRepositoryImpl struct to implement methods
type MoviesRepositoryImpl struct{}

// List method to list movies
func (r *MoviesRepositoryImpl) List() ([]entity.Movie, error) {
	res, err := http.Get("https://raw.githubusercontent.com/meilisearch/MeiliSearch/master/datasets/movies/movies.json")
	if err != nil {
		return []entity.Movie{}, errors.FailedGetResourceMovies
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []entity.Movie{}, errors.FailedToReadResponse
	}
	var movies = []entity.Movie{}
	json.Unmarshal(body, &movies)
	return movies, nil
}
