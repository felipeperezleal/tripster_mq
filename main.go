package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	rabbitmq "internal/rabbitmq.go/internal"
)

func main() {
	rabbitMQConn := rabbitmq.InitRabbitMQConnection()
	rabbitMQConsumer := rabbitmq.InitRabbitMQConsumer(rabbitMQConn)

	go rabbitMQConsumer.StartConsuming()

	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGINT, syscall.SIGTERM)
	<-sigTerm

	rabbitMQConn.Close()
	fmt.Println("Tripster MQ se ha detenido.")
}
