package main

import (
	"log/slog"
	"os"

	"github.com/ss49919201/go-functional-layerd-architecture/in-memory/server"
)

func main() {
	if err := server.ListenAndServe(8080); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
