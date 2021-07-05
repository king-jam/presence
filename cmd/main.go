package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/king-jam/presence/agent"
	"github.com/king-jam/presence/plugins/monitors"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var monitorPlugins []monitors.Plugin
	for name, initializer := range monitors.Registry {
		log.Println("Starting: ", name)
		m, err := initializer()
		if err != nil {
			log.Fatalf("failed to start %s", name)
		}
		monitorPlugins = append(monitorPlugins, m)
	}

	notifier := agent.NewNotifier()
	// Catch signal so we can shutdown gracefully
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Println("Starting")
		if err := monitorPlugins[0].Start(notifier); err != nil {
			log.Fatalf("Run Error: %s\n", err)
		}
	}()
	// defer will handle all the cleanup
	defer func() {
		err := monitorPlugins[0].Stop()
		if err != nil {
			log.Fatalf("Shutdown Error: %s\n", err)
		}
	}()

	// Wait for a signal before shutting down
	sig := <-sigCh
	log.Printf("%s Signal received. Shutting down\n", sig.String())
}
