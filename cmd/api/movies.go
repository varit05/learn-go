package main

import (
	"fmt"
	"net/http"
	"time"

	"learngo.varit.me/internal/data"
)

func (app *application) createMovieHandler(response http.ResponseWriter, r *http.Request) {
	fmt.Println(response, "create a new movie")
}

func (app *application) showMoviesHandler(showMoviesResponse http.ResponseWriter, showMoviesRequest *http.Request) {

	id, err := app.readIDParam(showMoviesRequest)

	if err != nil || id < 1 {
		http.NotFound(showMoviesResponse, showMoviesRequest)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Dark Night",
		Runtime:   102,
		Genres:    []string{"drama", "superhero"},
		Version:   1,
	}

	err = app.writeJSON(showMoviesResponse, http.StatusOK, envelope{"movie": movie}, nil)

	if err != nil {
		app.logger.Print(err)
		http.Error(showMoviesResponse, "Server could not process the response", http.StatusInternalServerError)
		return
	}
}
