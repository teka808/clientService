package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"test/internal/config"
)

type mq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func Init() *mq {
	mq := new(mq)
	mqConn, err := amqp.Dial(config.RABBITMQ_HOST)

	if err != nil {
		panic(err)
	}

	mq.conn = mqConn
	mq.OpenChannel()
	return mq
}

func (mq *mq) OpenChannel() {
	channelRabbitMQ, err := mq.conn.Channel()
	if err != nil {
		panic(err)
	}
	mq.channel = channelRabbitMQ

	log.Println("Successfully channel opened")
}

func (mq *mq) CloseChannel() {
	defer mq.channel.Close()
	defer mq.conn.Close()
}

func (mq *mq) Publish(message string) string {

	err := mq.channel.Publish(
		"",                            // exchange
		config.RABBITMQ_CONSUME_QUEUE, // routing key
		false,                         // mandatory
		false,                         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		return "fail"
	}
	return "success"
}
