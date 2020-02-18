package api

import (
	"fmt"
	"net/http"
	"strconv"


	"github.com/ejemba/fizzbuzz/application/leboncoin"
	"github.com/ejemba/fizzbuzz/domain/zz"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Gophers!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the fizzbuzz"))
}

func NewFizzBuzzHandler(logger log.Logger) *FizzBuzzHandler {
	return &FizzBuzzHandler{logger: logger}
}


type FizzBuzzHandler struct {
	logger log.Logger
}

func (fbHandler *FizzBuzzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//	http.DefaultServeMux
	logger := log.With(fbHandler.logger, "method", "ServeHTTP")
	
	request, err := createFizzBuzzRequest(r)
	level.Debug(logger).Log("request", fmt.Sprintf("%#v", request))

	if err != nil {
		handleError(w, logger, err)
	} else {
		err = request.Validate()
		if err != nil {
			handleError(w, logger, err)
		} else {
			z1 := zz.NewZz(request.Int1, request.Str1)
			z2 := zz.NewZz(request.Int2, request.Str2)
			limit := request.Limit

			lbcService := leboncoin.NewLeBonCoinService(z1, z2, limit)
			response := lbcService.Execute()

			RootTopRequests.AddStat(request)

			level.Info(logger).Log("msg", response)

			w.Write([]byte(response))
		}
	}
}

func handleError(w http.ResponseWriter, logger log.Logger, e error) {
	level.Error(logger).Log("msg", e.Error(), "status", http.StatusBadRequest)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(e.Error() + "\n"))
}

func requestInspection(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(string("Method: " + r.Method + "\n")))
	w.Write([]byte(string("Protocol Version: " + r.Proto + "\n")))
	w.Write([]byte(string("Host: " + r.Host + "\n")))
	w.Write([]byte(string("Path: " + r.URL.Path + "\n")))
	w.Write([]byte(string("Referer: " + r.Referer() + "\n")))
	w.Write([]byte(string("User Agent: " + r.UserAgent() + "\n")))
	w.Write([]byte(string("Remote Addr: " + r.RemoteAddr + "\n")))
	w.Write([]byte(string("Requested URL: " + r.RequestURI + "\n")))
	w.Write([]byte(string("Content Length: " + strconv.FormatInt(r.ContentLength, 10) + "\n")))
	for key, value := range r.URL.Query() {
		w.Write([]byte(string("Query string: key=" + key + " value=" + value[0] + "\n")))
	}

}
