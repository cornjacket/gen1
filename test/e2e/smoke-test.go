package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883") // MQTT broker address

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	topic := "direct/david"
	message := "Hello, MQTT!"

	for i:=0; i<5; i++ {
		token := client.Publish(topic, 0, false, message)
		token.Wait()

		fmt.Printf("Test message published to topic %s: %s\n", topic, message)

		// Wait for a few seconds to receive messages
		time.Sleep(5 * time.Second)
	}

	// Disconnect from the broker
	client.Disconnect(250)
	fmt.Println("MQTT client has been disconnected.")
}


