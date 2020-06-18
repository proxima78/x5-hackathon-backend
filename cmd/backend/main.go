package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/proxima78/x5-hackathon-backend/app"
	"github.com/proxima78/x5-hackathon-backend/config"
	"github.com/proxima78/x5-hackathon-backend/router"
)

func main() {
	var configPath string
	var generateConfig bool
	flag.StringVar(&configPath, "config", "", "Path to server config")
	flag.BoolVar(&generateConfig, "genconfig", false, "Generate new config")
	flag.Parse()
	if generateConfig {
		confStr, err := config.Generate()
		if err != nil {
			log.Fatalf("Cannot generate config! %s", err.Error())
		}
		fmt.Print(confStr)
		os.Exit(0)
	}
	log.Print("Starting backend...")
	if configPath == "" {
		log.Fatal("Path to config isn't specified!")
	}

	cfg, err := config.Parse(configPath)
	if err != nil {
		log.Fatal(err)
	}
	app, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}
	router, err := router.NewRouter(app)
	if err != nil {
		log.Fatalf("Failed to initialize router: %s", err.Error())
	}

	// CTRL+C handler.
	signalHandler := make(chan os.Signal, 1)
	shutdownDone := make(chan bool, 1)
	signal.Notify(signalHandler, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalHandler
		log.Println("CTRL+C or SIGTERM received, shutting down backend...")
		app.Destroy()
		shutdownDone <- true
	}()

	app.Run(router)

	<-shutdownDone
	os.Exit(0)
}
