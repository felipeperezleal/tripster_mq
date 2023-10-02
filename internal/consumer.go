package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	conn *amqp.Connection
}

func InitRabbitMQConsumer(conn *RabbitMQConnection) *RabbitMQConsumer {
	return &RabbitMQConsumer{conn: conn.Conn}
}

func (c *RabbitMQConsumer) StartConsuming() {
	ch, err := c.conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	queueName := "tripster_queue"

	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Esperando mensajes de RabbitMQ...")

	for msg := range msgs {
		fmt.Printf("Recibido mensaje: %s\n", msg.Body)
	}
}
