package main

import (
	"fmt"

	usersapp "github.com/wrtgvr/grpc-microblog/internal/app/users"
)

func main() {
	app := usersapp.NewApp(50051)
	app.MustRun()

	fmt.Println("Server running")
}
