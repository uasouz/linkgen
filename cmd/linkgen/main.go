package main

import (
	"context"
	"fmt"
	"linkgen/config"
	"linkgen/linkgen"
	"linkgen/store"
	"linkgen/store/memory"
	"linkgen/store/mysql"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func initializeLinkStore(cfg config.LinkStoreConfig) (store.LinkStore, error) {
	switch cfg.Type {
	case "mysql":
		return mysql.New(cfg.DSN)
	default:
		return memory.New(), nil
	}
	return nil, nil
}

func main() {
	// Start configuration structure with default values
	cfg := config.Config{
		APIPort: "3000",
	}

	// Call configuration load function
	cfg.LoadConfig("config.yaml")

	// Get new instance of store.LinkStore
	linkStore, err := initializeLinkStore(cfg.LinkStoreConfig)

	// Panic if there is no storage
	if err != nil {
		panic("store failed to initialize")
	}

	// Creates a new API Service with the port configuration
	linkgenService := linkgen.New(cfg.APIPort, linkStore)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	g, errGroupCtx := errgroup.WithContext(ctx)

	// Start API Service
	g.Go(linkgenService.Start)

	g.Go(func() error {
		<-errGroupCtx.Done()

		//// Create a new context to shutdown gracefully
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		return linkgenService.Stop(shutdownCtx)
	})

	// Wait for all goroutines to finish
	if err = g.Wait(); err == context.Canceled || err == nil {
		fmt.Println("gracefully quit server")
	} else if err != nil {
		fmt.Println(err)
		fmt.Println("server quit unexpectedly")
	}

}
