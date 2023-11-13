# Tripster MQ

Tripster MQ es un componente que utiliza RabbitMQ para proporcionar una cola de mensajes. Este microservicio se encarga de recibir mensajes de otros componentes de la aplicación y procesarlos según sea necesario.

## Requisitos

- Docker (para ejecutar RabbitMQ)
- Asegúrate de haber descargado la imagen de docker anteriormente
  ```cmd
   docker pull felipeperezleal/tripster_mq
  ```

## Ejecución

Para ejecutar Tripster MQ, primero asegúrate de tener Docker instalado. Luego, puedes seguir estos pasos:

1. Ejecuta el siguiente comando para iniciar un contenedor Docker de RabbitMQ:

   ```bash
   docker run 5672:5672 felipeperezleal/tripster_mq
   ```

Esto iniciará un contenedor RabbitMQ que Tripster MQ utilizará para gestionar la cola de mensajes.

Tripster MQ estará listo para recibir y procesar mensajes de otros componentes de la aplicación.

## Uso
Tripster MQ escuchará la cola de mensajes y procesará los mensajes entrantes.
