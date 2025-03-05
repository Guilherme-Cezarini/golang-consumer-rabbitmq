package delivery

import (
	"consumer-golang/internal/domain"
	"consumer-golang/internal/service"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQConsumer struct {
	conn           *amqp.Connection
	channel        *amqp.Channel
	messageService *service.MessageService
	queueName      string
}

func NewRabbitMQConsumer(uri string, messageService *service.MessageService, queueName string) (*RabbitMQConsumer, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQConsumer{
		conn:           conn,
		channel:        channel,
		messageService: messageService,
		queueName:      queueName,
	}, nil
}

func (c *RabbitMQConsumer) Consume() {
	queue, err := c.channel.QueueDeclare(
		c.queueName, // queue name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := c.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var message domain.Message
			err := json.Unmarshal(d.Body, &message)
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}

			err = c.messageService.ProcessMessage(&message)
			if err != nil {
				log.Printf("Error processing message: %v", err)
			} else {
				log.Printf("Message processed: %v", message)
			}
		}
	}()

	log.Printf("Waiting for messages...")
	<-forever
}
