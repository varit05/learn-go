package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(response http.ResponseWriter, r *http.Request) {
	fmt.Println(response, "create a new movie")
}

func (app *application) showMoviesHandler(showMoviesResponse http.ResponseWriter, showMoviesRequest *http.Request) {

	id, err := app.readIDParam(showMoviesRequest)

	if err != nil || id < 1 {
		http.NotFound(showMoviesResponse, showMoviesRequest)
	}
	fmt.Fprintf(showMoviesResponse, "Show the details of movie %d\n", id)
}
