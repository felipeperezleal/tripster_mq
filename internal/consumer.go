package rabbitmq

import (
	"fmt"
	"log"
	"time"

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
	} else {
		fmt.Printf("Cola %s creada exitosamente.\n", queueName)
	}

	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Esperando mensajes de RabbitMQ...")

	for msg := range msgs {
		fmt.Printf("Recibido mensaje: %s\n", msg.Body)

		fmt.Println("Esperando a que el algoritmo se complete...")
		time.Sleep(5 * time.Second)
		fmt.Println("El algoritmo se ha completado.")
	}
}
