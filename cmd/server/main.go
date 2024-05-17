package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"os/signal"
)

func main() {
	connString := "amqp://guest:guest@localhost:5672/"
	connection, err := amqp.Dial(connString)
	if !(err == nil) {
		fmt.Println("error")
	} else {
		fmt.Println("Starting Peril server...")
	}
	defer connection.Close()
	// wait for ctrl+c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("Ctrl+C pressed. Exiting...")
}
