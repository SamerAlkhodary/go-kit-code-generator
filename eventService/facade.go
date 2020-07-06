package eventService

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"crypto/tls"
)

func Serve() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{

		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "eventService",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")
	var db *sql.DB
	{
		var err error
		db, err = sql.Open("mysql", "username:password@tcp(localhost)/db")
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	repository := MakeNewRepository(db, logger)

	flag.Parse()
	ctx := context.Background()
	var service EventService
	{
		service = NewService(logger, repository)
	}

	errs := make(chan error)
	go func() {

		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	endpoints := MakeEndpoints(service)
	go func() {
		fmt.Println("Listening on port", *httpAddr)
		//TODO: fill cert and key names
		cert, error := tls.LoadX509KeyPair("../keys/*.crt", "../keys/*.key")
		errs <- error
		handler := NewHTTPServer(ctx, endpoints)

		server := &http.Server{

			Addr:    *httpAddr,
			Handler: handler,
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{cert}},
		}
		errs <- server.ListenAndServeTLS("", "")
	}()
	level.Error(logger).Log("exit", <-errs)
}
