package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Устанавливаем соединение с сервером RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Ошибка при установке соединения с RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Создаем канал
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Ошибка при создании канала: %s", err)
	}
	defer ch.Close()

	// Объявляем очередь
	q, err := ch.QueueDeclare(
		"hello", // Имя очереди
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("Ошибка при объявлении очереди: %s", err)
	}

	// Получаем сообщения из очереди
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Ошибка при получении сообщений из очереди: %s", err)
	}

	// Обрабатываем полученные сообщения
	for msg := range msgs {
		log.Printf("Получено сообщение: %s", msg.Body)
	}
}
