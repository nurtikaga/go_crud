package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthCheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/createModule", app.createClotheHandler)
	router.HandlerFunc(http.MethodGet, "/v1/getModule/:id", app.getClotheHandler)
	router.HandlerFunc(http.MethodGet, "/v1/getModule", app.listClotheHandler)
	router.HandlerFunc(http.MethodPut, "/v1/updateModule/:id", app.editClotheHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/deleteModule/:id", app.deleteClotheHandler)

	return router
}
