package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Emmiter struct {
	conn *amqp.Connection
}

func (e *Emmiter) setup() error {
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	return declareExchange(channel)
}

func (e *Emmiter) Push(event string, severity string) error {
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	log.Println("Pushing event: ", event, " with severity: ", severity)

	err = channel.Publish(
		"logs_topic", // exchange
		severity,     // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func NewEventEmitter(conn *amqp.Connection) (Emmiter, error) {
	emmiter := Emmiter{
		conn: conn,
	}
	err := emmiter.setup()
	if err != nil {
		return Emmiter{}, err
	}
	return emmiter, nil
}
