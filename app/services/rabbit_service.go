package services

import (
	"Walker/global"
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitService struct {
}

func (service *RabbitService) Conn() (amqp.Queue, *amqp.Channel) {
	str := fmt.Sprintf("amqp://%s:%s@rabbitmq:5672/", "guest", "guest")
	conn, err := amqp.Dial(str)
	if err != nil {
		global.Logger.Error("连接rabbit发生错误" + err.Error())
	}
	channel, err := conn.Channel()
	if err != nil {
		global.Logger.Error("管道创建失败" + err.Error())
	}
	queue, err := channel.QueueDeclare("hello", true, false, false, false, nil)
	if err != nil {
		global.Logger.Error("创建队列失败" + err.Error())
	}
	return queue, channel
}
