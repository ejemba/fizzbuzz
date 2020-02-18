package main

import (
	"os"
	"github.com/ejemba/fizzbuzz/interfaces/api"
	"net/http"
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
	
	http.ListenAndServe(":3000", api.Mux) 
}
