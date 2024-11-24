package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	service "github.com/sergereinov/go-windows-service"
)

var (
	// You can set the Version at compile stage of your dev pipeline with:
	// go build -ldflags="-X main.Version=1.0.0" ./example
	Version     = "0.1"
	Name        = service.ExecutableFilename()
	Description = "My service"
)

func main() {
	logger := log.Default()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
	}

	file, err := os.OpenFile(fmt.Sprintf("%s\\logs\\%s-%s.log", dir, Name, time.Now().Format("2006-01-02")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer file.Close()

	logger.SetOutput(file)

	// Run service wrapper
	service.Service{
		Version:     Version,
		Name:        Name,
		Description: Description,
		Logger:      logger,
	}.Proceed(func(ctx context.Context) {

		logger.Printf("Service %s v%s started", Name, Version)

		if err := openBrowser(); err != nil {
			logger.Printf("Error changing the wallpaper %s", err)
		} else {
			logger.Printf("Opened browser")
		}

		<-ctx.Done()

		if err := changeWallpaper(); err != nil {
			logger.Printf("Error changing the wallpaper %s", err)
		} else {
			logger.Printf("Changed wallpapers")
		}

		logger.Printf("Service %s v%s stopped", Name, Version)
	})
}
