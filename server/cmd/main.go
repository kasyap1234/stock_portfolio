package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/knadh/koanf/v2"
	"github.com/labstack/echo"
)

var k = koanf.New(".")

func main() {
	e := echo.New()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		e.Start(":" + k.Int(server.port))

	}()
	<-quit
	shutdownCntx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("shutting down")
	if err := e.Shutdown(shutdownCntx); err != nil {
		log.Fatal("Server failed to shutdown gracefully")
	}
	log.Print("server stopped")

}
