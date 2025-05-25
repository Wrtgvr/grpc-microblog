package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	usersapp "github.com/wrtgvr/grpc-microblog/internal/app/users"
)

var (
	stop = make(chan os.Signal, 1)
)

func main() {
	app := usersapp.NewApp(50051)
	go app.MustRun()

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	fmt.Println("Stopping server")
	app.Stop()
	fmt.Println("Server stopped")
}
