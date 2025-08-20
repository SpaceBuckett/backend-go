package main

import (
	"fmt"
	"log/slog"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var counter int
	fmt.Printf("Hello, Shahzad!", counter)
}
