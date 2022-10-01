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

func Init() {
	mq := new(mq)
	mqConn, err := amqp.Dial(config.RABBITMQ_HOST)

	if err != nil {
		panic(err)
	}

	mq.conn = mqConn

	mq.OpenChannel()
	mq.Consume()
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

func (mq *mq) Consume() {
	//messages, err := mq.channel.Consume(
	//	config.RABBITMQ_CONSUME_QUEUE, // queue name
	//	"",                            // consumer
	//	true,                          // auto-ack
	//	false,                         // exclusive
	//	false,                         // no local
	//	false,                         // no wait
	//	nil,                           // arguments
	//)
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("Successfully connected to RabbitMQ")

	//go func() {
	//	for {
	//		select {
	//		case data := <-messages:
	//
	//			var info map[string]interface{}
	//			e := json.Unmarshal(data.Body, &info)
	//
	//			if e != nil {
	//				panic(e)
	//			}
	//
	//			for _, v := range info {
	//				coordinates := v.(map[string]interface{})["coordinates"]
	//				latitude := 0.0
	//				longitude := 0.0
	//				if coordinates != nil {
	//					latitude = v.(map[string]interface{})["coordinates"].([]interface{})[0].(float64)
	//					longitude = v.(map[string]interface{})["coordinates"].([]interface{})[1].(float64)
	//				}
	//				db.InsertOrUpdate(
	//					v.(map[string]interface{})["id"].(string),
	//					v.(map[string]interface{})["type"].(string),
	//					latitude,
	//					longitude,
	//					v.(map[string]interface{})["status"].(string),
	//					v.(map[string]interface{})["timezone"].(string),
	//				)
	//				log.Println("Inserted/Updated data")
	//			}
	//		}
	//	}
	//}()
}
