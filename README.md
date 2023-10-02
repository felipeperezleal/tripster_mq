# Tripster MQ

Tripster MQ es un microservicio que utiliza RabbitMQ para proporcionar una cola de mensajes. Este microservicio se encarga de recibir mensajes de otros componentes de la aplicación y procesarlos según sea necesario.

## Requisitos

- Docker (para ejecutar RabbitMQ)

## Ejecución

Para ejecutar Tripster MQ, primero asegúrate de tener Docker instalado. Luego, puedes seguir estos pasos:

1. Ejecuta el siguiente comando para iniciar un contenedor Docker de RabbitMQ:

   ```bash
   docker run -d --name tripster_mq -p 5672:5672 rabbitmq:3
   ```

Esto iniciará un contenedor RabbitMQ que Tripster MQ utilizará para gestionar la cola de mensajes.

### Compilar y ejecutar Tripster MQ:

  ```bash
  go build -o tripster_mq main.go
  ./tripster_mq
  ```
Esto compilará el código de Tripster MQ y lo ejecutará. El microservicio comenzará a escuchar la cola de mensajes de RabbitMQ y procesará los mensajes entrantes según sea necesario.

Tripster MQ estará listo para recibir y procesar mensajes de otros componentes de la aplicación.

## Uso
Para utilizar Tripster MQ en tu aplicación, puedes enviar mensajes a la cola de RabbitMQ utilizando una biblioteca de RabbitMQ compatible con Go. Tripster MQ escuchará la cola de mensajes y procesará los mensajes entrantes.
