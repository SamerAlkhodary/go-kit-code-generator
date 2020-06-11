package chatService

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func main() {
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
	flag.Pars()
	ctx := context.Background()
	var service chatService.ChatService
	{
		service = chatService.NewServer(logger)
	}

	errs := make(chan error)
	go func() {

		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	endpoints := chatService.MakeEndpoints(service)
	go func() {
		fmt.Println("Listening on port", *httpAddr)
		handler := chatService.NewHTTPServer(ctx, endpoints)

	}()
	level.Error(looger).Log("exit", <-errs)
}
