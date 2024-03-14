package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createMovieHandler(response http.ResponseWriter, r *http.Request) {
	fmt.Println(response, "create a new movie")
}

func (app *application) showMoviesHandler(showMoviesResponse http.ResponseWriter, showMoviesRequest *http.Request) {

	params := httprouter.ParamsFromContext(showMoviesRequest.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		http.NotFound(showMoviesResponse, showMoviesRequest)
	}
	fmt.Fprintln(showMoviesResponse, "Show the details of movie %d\n", id)
}
