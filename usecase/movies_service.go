package usecase

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	pb "github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"github.com/thalysonalexr/movie-poster/entity"
	"github.com/thalysonalexr/movie-poster/infra/repo"
)

// Service interface of usecase
type Service interface {
	SearchByGender(k string) ([]entity.Movie, error)
	DownloadPosters(k string) (bool, error)
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

func downloadFile(URL string, fileName string, c chan string) (string, error) {
	defer close(c)
	res, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", errors.New("Received non 200 response code")
	}
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return "", err
	}
	c <- URL
	return URL, nil
}

func makeChannels(movies []entity.Movie, cRes chan string) {
	defer close(cRes)
	var results = []chan string{}
	path, _ := os.Getwd()

	for i := range movies {
		filePath := filepath.FromSlash(path + "/../tmp/" + renameFile(movies[i].Poster))
		results = append(results, make(chan string))
		go downloadFile(movies[i].Poster, filePath, results[i])
	}
	for i := range results {
		for r := range results[i] {
			cRes <- r
		}
	}
}

func renameFile(name string) string {
	splited := strings.Split(strings.ReplaceAll(name, " ", "-"), "/")
	return strings.ToLower(splited[len(splited)-1])
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

// DownloadPosters download poster images
func (service *ServiceImpl) DownloadPosters(k string) (bool, error) {
	movies, err := service.SearchByGender(k)
	if err != nil {
		return false, err
	}
	ini := time.Now()
	results := make(chan string)
	go makeChannels(movies, results)

	bar := pb.Full.Start(len(movies))

	for range results {
		bar.Increment()
	}

	bar.Finish()
	color.Cyan("Service finish proccess. Duration %f secs\n", time.Since(ini).Seconds())
	return true, nil
}
