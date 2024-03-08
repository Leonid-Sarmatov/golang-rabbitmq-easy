package main

import (
	"log"
	"time"

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

	// Отправляем сообщения в очередь
	for {
		message := "Hello, RabbitMQ!"
		err := ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			})
		if err != nil {
			log.Printf("Ошибка при отправке сообщения: %s", err)
		} else {
			log.Printf("Отправлено сообщение: %s", message)
		}
		time.Sleep(5 * time.Second) // Период отправки сообщений
	}
}
