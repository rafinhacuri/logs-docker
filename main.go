package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	if err := initTelegram(); err != nil {
		log.Fatal("Erro ao iniciar Telegram:", err)
	}
	if err := initDiscord(); err != nil {
		log.Fatal("Erro ao iniciar Discord:", err)
	}
	if err := initDocker(); err != nil {
		log.Fatal("Erro ao iniciar Docker:", err)
	}
}

func main() {
	defer discordSession.Close()
	defer clientDocker.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go followNewContainers(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
