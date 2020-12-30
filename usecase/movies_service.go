package usecase

import (
	"strings"

	"github.com/thalysonalexr/movie-poster/entity"
	"github.com/thalysonalexr/movie-poster/infra/repo"
)

// Service interface of usecase
type Service interface {
	SearchByGender(k string) ([]entity.Movie, error)
}

// ServiceImpl is implementation of Service interface
type ServiceImpl struct {
	repo repo.Repository
}

// CreateNewService is a factory to create new Service
func CreateNewService(r repo.Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: r,
	}
}

// SearchByGender is a handler to get list movies
func (service *ServiceImpl) SearchByGender(k string) ([]entity.Movie, error) {
	movies, err := service.repo.List()
	if err != nil {
		return []entity.Movie{}, err
	}
	filtered := []entity.Movie{}
	for i := range movies {
		for j := range movies[i].Genres {
			if strings.Contains(strings.ToLower(movies[i].Genres[j]), strings.ToLower(k)) {
				filtered = append(filtered, movies[i])
				break
			}
		}
	}
	return filtered, nil
}
