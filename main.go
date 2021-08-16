package main

import (
	"fmt"
	"net/http"
	"time"

	nats "github.com/nats-io/nats.go"

	"github.com/lenvendo/ig-absolut-fake-sms/lib/container"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
	"os"
	"context"
	"errors"
)

const (
	GracefulShutdownTimeout = 5 * time.Second
	AppShutdownMessage      = "program exits"
)

func main() {
	isShuttingDown := false
	errs := make(chan error)

	c := container.NewContainer()
	cfg := c.Config

	nc, err := nats.Connect(
		fmt.Sprintf("nats://%s:%d", cfg.Nats.Host, cfg.Nats.Port),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(cfg.Nats.RetryLimit),
		nats.ReconnectWait(time.Millisecond*time.Duration(cfg.Nats.WaitLimit)),
		nats.UserInfo(cfg.Nats.Username, cfg.Nats.Password),
	)
	if err != nil {
		c.Logger.Error("failed to connect nats", zap.Error(err))
		return
	}
	defer nc.Close()

	c.Logger.Info("Connected to NATS at:", zap.String("ConnectedUrl", nc.ConnectedUrl()))
	subscription, err := nc.Subscribe("tasks", c.WorkerService.SendHandler)

	var codesHttpServer *http.Server
	{
		m := http.NewServeMux()
		m.Handle("/codes", c.CodesHandler)
		var handler http.Handler = m

		codesHttpServer = &http.Server{
			Addr:         ":" + cfg.HTTP.Port,
			ReadTimeout:  60 * time.Second,
			WriteTimeout: 60 * time.Second,
			Handler:      handler,
		}
	}

	go func() {
		c.Logger.Info("starting Codes HTTP server", zap.String("addr", codesHttpServer.Addr))
		errs <- codesHttpServer.ListenAndServe()
	}()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		sig := <-quit
		c.Logger.Warn("OS Signal recieved, starting graceful shutdown with timeout",
			zap.String("signal", sig.String()),
			zap.Duration("timeout", GracefulShutdownTimeout),
		)
		isShuttingDown = true

		success := make(chan string)
		go func() {
			errs <- codesHttpServer.Shutdown(context.Background())
			success <- "codesHttpServer successfully stopped"
		}()

		go func() {
			errs <- subscription.Unsubscribe()
			success <- "subscription successfully cancelled"
		}()

		go func() {
			time.Sleep(GracefulShutdownTimeout)
			c.Logger.Fatal("shutting down timeout reached, stopping through Fatal")
		}()

		c.Logger.Info(<-success)
		c.Logger.Info(<-success)
		errs <- errors.New(AppShutdownMessage)
	}()

	for err := range errs {
		if err == nil {
			continue
		}
		if !isShuttingDown {
			c.Logger.Fatal("fatal shutdown", zap.Error(err))
		}
		c.Logger.Warn("shutdown err message", zap.Error(err))

		if err.Error() == AppShutdownMessage {
			return
		}
	}
}
