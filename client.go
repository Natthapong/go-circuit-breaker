package main

import (
	"go-circuit-breaker-server/circuit"
	"log"
	"time"
)

func main() {
	var client circuit.NotificationClient
	// Create a common Client
	client = circuit.NewSmsClient("http://localhost:8080")
	// And then wrap it
	client = circuit.NewClientCircuitBreakerProxy(client)

	for {
		err := client.Send()
		time.Sleep(1 * time.Second)
		if err != nil {
			log.Println("caught an error", err)
		}
	}
}
