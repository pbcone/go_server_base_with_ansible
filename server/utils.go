package main

import (
	"fmt"
	"net/http"
	"os"

	"./config"
	"github.com/julienschmidt/httprouter"
)

//Index heartbeat functionality
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "go Server Version:  "+config.Version+"\n")
}

//EnableCors this will enable CORS
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func padNumberWithZero(value int) string {
	return fmt.Sprintf("%06d", value)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
