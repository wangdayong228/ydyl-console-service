package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wangdayong228/ydyl-console-service/internal/configs"
	"github.com/wangdayong228/ydyl-console-service/internal/servers"
)

func init() {
	configs.Init()
	servers.Init()
}

func main() {
	// services.Start()
	_servers := servers.Start()
	defer (func() {
		servers.Stop(_servers)
	})()

	waitQuitSignal()
	log.Println("Server exiting")

}

func waitQuitSignal() {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down servers...")
}
