package chatService

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
)

func Serve() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{

		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "chatService",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")
	var db *sql.DB
	{
		var err error
		db, err = sql.Open("mysql", "address")
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	repository := MakeNewRepository(db, logger)

	flag.Parse()
	ctx := context.Background()
	var service ChatService
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
		handler := NewHTTPServer(ctx, endpoints)

		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	level.Error(logger).Log("exit", <-errs)
}
