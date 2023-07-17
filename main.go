package main

import (
	"widget-sports/configurations"
	"widget-sports/handlers"
)

func main() {
	cfg := configurations.LoadConfig()
	handlers.Handlers(*cfg)
}
