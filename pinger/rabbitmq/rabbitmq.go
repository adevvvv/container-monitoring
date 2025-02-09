package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"container-monitoring/pinger/api"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
	Queue   amqp091.Queue
}

func NewRabbitMQ() (*RabbitMQ, error) {
	rabbitURL := "amqp://guest:guest@rabbitmq:5672/"
	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create RabbitMQ channel: %v", err)
	}

	q, err := ch.QueueDeclare(
		"ping_queue", 
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %v", err)
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
		Queue:   q,
	}, nil
}

func (rmq *RabbitMQ) SendMessage(status api.PingStatus) error {
	data, err := json.Marshal(status)
	if err != nil {
		return fmt.Errorf("error marshaling: %v", err)
	}

	err = rmq.Channel.Publish(
		"",
		rmq.Queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
	if err != nil {
		return fmt.Errorf("error sending message to RabbitMQ: %v", err)
	}
	log.Printf("Sent status to RabbitMQ: %+v", status)
	return nil
}

func (rmq *RabbitMQ) Close() {
	if rmq.Channel != nil {
		rmq.Channel.Close()
	}
	if rmq.Conn != nil {
		rmq.Conn.Close()
	}
}
