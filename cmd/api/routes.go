package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "v1/healthCheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "v1/movie", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "v1/movies/:id", app.showMoviesHandler)

	return router
}
