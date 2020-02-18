package api

import (
	"net/http"
	"github.com/go-kit/kit/log"
	"expvar"
)


var Mux = http.NewServeMux()

func RegisterRoutes(logger log.Logger) {	
	Mux.Handle("/fizzbuzz", NewFizzBuzzHandler(logger))
	Mux.Handle("/debug/vars", expvar.Handler()) 
}
