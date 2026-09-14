package main

import (
	"fmt"

	"github.com/hekemen/automata/internal/infrastructure/config"
)

func main() {
	config.Load("config.yaml")

	fmt.Println("automata health: ok")
	fmt.Printf("  server.host = %s\n", config.Get("server.host"))
	fmt.Printf("  server.port = %s\n", config.Get("server.port"))
}
