package consumer

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
)

func InitializeConnectionChannelRabbitMQ() {
	InitializeConsumer()
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	queueName := "orders"
	// _, err = ch.QueueDeclare(
	// 	queueName, // queue
	// 	true,      // durable
	// 	false,     // auto-delete
	// 	false,     // exclusive
	// 	false,     // no-wait
	// 	nil,       // args
	// )
	// if err != nil {
	// 	log.Printf("Failed to declare a queue: %v", err)
	// 	return err
	// }

	msgs, err = ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			consumeMessage(d)
		}
	}()

	log.Println("RabbitMQ consumer initialized and listening for messages")

	// Aguarda um sinal para encerrar (CTRL+C)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	// Aguarda o sinal de encerramento
	<-sig

	log.Println("Shutting down RabbitMQ consumer")
}
