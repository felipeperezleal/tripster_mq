package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConnection struct {
	Conn *amqp.Connection
}

func InitRabbitMQConnection() *RabbitMQConnection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión a RabbitMQ establecida.")
	return &RabbitMQConnection{Conn: conn}
}
