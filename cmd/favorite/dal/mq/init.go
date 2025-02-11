package mq

import (
	"context"
	"douyin/cmd/favorite/dal/mq/amqpclt"
	"fmt"

	"douyin/pkg/constants"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

var AddActor *amqpclt.Actor

// Init InitMq to init rabbitMQ
func Init() {
	amqpConn, err := amqp.Dial(fmt.Sprintf(constants.RabbitMqURI,
		constants.MQUser, constants.MQPassword, constants.MQHost, constants.MQPort))
	if err != nil {
		klog.Fatal("cannot dial amqp", err)
	}
	AddActor, err = amqpclt.NewActor(amqpConn, "favorite_action")
	if err != nil {
		klog.Fatal("cannot create add actor", err)
	}

	//开启消费监听
	go AddActor.Consumer(context.Background())
}
