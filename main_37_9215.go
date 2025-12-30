package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("[INIT] Starting MQTT Client...")
	broker := "tcp://iot-broker.base58labs.internal:1883"
	
	for {
		// Simulate connection and heartbeat
		fmt.Printf("[CONN] Connected to broker: %s\n", broker)
		fmt.Println("[PUB] Topic: sensors/data | Payload: {temp: 24.5, humidity: 60}")
		
		time.Sleep(60 * time.Second)
	}
}
