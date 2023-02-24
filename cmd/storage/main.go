package main

import (
	"fmt"

	"github.com/semerf/WeatherServer/internal/server"
)

func main() {
	fmt.Println("Hello, world!")
	go server.Server()
}
