package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/felipeperezleal/tripster_mq/internal/rabbitmq"
)

func main() {
	// Inicializa la conexión a RabbitMQ
	rabbitMQConn := rabbitmq.InitRabbitMQConnection()

	// Inicializa el consumidor de RabbitMQ
	rabbitMQConsumer := rabbitmq.InitRabbitMQConsumer(rabbitMQConn)

	// Inicia el consumidor
	go rabbitMQConsumer.StartConsuming()

	// Maneja señales de terminación
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGINT, syscall.SIGTERM)
	<-sigTerm

	// Cierra la conexión de RabbitMQ y otros recursos
	rabbitMQConn.Close()
	fmt.Println("Tripster MQ se ha detenido.")
}
