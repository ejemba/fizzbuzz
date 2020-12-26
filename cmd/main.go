package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ejemba/fizzbuzz/interfaces/api"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func init() {
	api.RegisterMetrics()
}

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "fizzbuzz",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	api.RegisterRoutes(logger)

	ip := os.Getenv("IP")
	port := os.Getenv("PORT")

	if ip == "" {
		ip = "::"
	}

	if port == "" {
		port = "3000"
	}

	addr := fmt.Sprintf("[%s]:%s", ip, port)

	level.Info(logger).Log("msg", strings.Join([]string{"listening to ", addr}, ""))

	http.ListenAndServe(addr, api.Mux)
}
