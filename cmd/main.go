package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/arturzhamaliyev/p2p-transfer/internal/controller"
	"github.com/arturzhamaliyev/p2p-transfer/internal/db"
	"github.com/arturzhamaliyev/p2p-transfer/internal/repository"
	"github.com/arturzhamaliyev/p2p-transfer/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func main() {
	cfg := koanf.New(".")
	err := cfg.Load(file.Provider("config.json"), json.Parser())
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// connect to db
	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close(ctx)

	// repo
	repo := repository.NewRepository(cfg, database)

	// service
	service := service.NewService(cfg, repo)

	// handler
	router := gin.Default()
	handler := controller.NewHandler(cfg, router, service)
	handler.InitRoutes()

	out := make(chan os.Signal, 1)
	signal.Notify(out, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-out
		fmt.Println("app is shutting down gracefully")
		database.Close(ctx)
		os.Exit(0)
	}()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Int("port")), router); err != nil {
		log.Fatal(err)
	}
}
