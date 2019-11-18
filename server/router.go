package main

import (
	"github.com/julienschmidt/httprouter"
)

func serverRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", index)
	return router
}
