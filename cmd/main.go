package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/fatih/color"

	repository "github.com/thalysonalexr/movie-poster/infra/repo"
	service "github.com/thalysonalexr/movie-poster/usecase"
)

var InvalidNumberOfArguments = errors.New("Invalid number of arguments")
var FlagNotExists = errors.New("Flag not exists. Pass valid flags, to show --help")

const (
	maxArgs = 3
	minArgs = 2
)

// Command struct to commands
type Command struct {
	Name        string
	Description string
}

var commands = []Command{
	{
		Name:        "help",
		Description: "Show list of available commmands",
	},
	{
		Name:        "gender",
		Description: "Search movies by gender",
	},
	{
		Name:        "poster-download",
		Description: "Download all posters searching by gender",
	},
}

func readFlags(command Command, exec func(k ...string)) error {
	if len(os.Args) < minArgs || len(os.Args) > maxArgs {
		return InvalidNumberOfArguments
	}
	f := os.Args[1]
	if strings.ToLower(f) == "--"+command.Name {
		var v string
		if len(os.Args) == 3 {
			v = os.Args[2]
		}

		exec(v)
		os.Exit(0)
	}
	stay := false
	for i := range commands {
		if strings.ToLower(f) == "--"+commands[i].Name {
			stay = true
			break
		}
	}
	if !stay {
		return FlagNotExists
	}
	return nil
}

func showHelper(v ...string) {
	color.Cyan("List of valid commands:")
	for i := range commands {
		color.Cyan("  --" + commands[i].Name + "     - " + commands[i].Description + "\n")
	}
}

func main() {
	err := readFlags(commands[0], showHelper)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	repo := repository.MoviesRepositoryImpl{}
	service := service.CreateNewService(&repo)

	err = readFlags(commands[1], func(keyword ...string) {
		movies, err := service.SearchByGender(keyword[0])
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		for i := range movies {
			var prettyJSON bytes.Buffer
			data, _ := json.Marshal(movies[i])
			error := json.Indent(&prettyJSON, data, "", "  ")
			if error != nil {
				color.Red("JSON parse error: ")
				os.Exit(1)
			}
			color.Yellow(string(prettyJSON.Bytes()))
		}
	})
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	err = readFlags(commands[2], func(keyword ...string) {
		success, err := service.DownloadPosters(keyword[0])
		if err != nil || !success {
			color.Red(err.Error())
			os.Exit(1)
		}
	})
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}
