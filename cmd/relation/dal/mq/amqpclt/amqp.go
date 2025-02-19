package amqpclt

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
)

// Actor implements an amqp Actor.
type Actor struct {
	channel   *amqp.Channel
	exchange  string // 不指定，走默认路由
	queueName string
}

func NewActor(conn *amqp.Connection, queueName string) (*Actor, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}
	return &Actor{
		channel:   ch,
		queueName: queueName,
	}, nil
}

// Publish 生产消息
func (a *Actor) Publish(_ context.Context, relation string) error {
	// 队列名 持久化 自动删除 exclusive 是否阻塞 额外参数
	_, err := a.channel.QueueDeclare(a.queueName, false,
		false, false, false, nil)
	if err != nil {
		return fmt.Errorf("cannot declare queue: %v", err)
	}

	// 投递消息
	return a.channel.Publish(a.exchange, a.queueName,
		false, false, amqp.Publishing{
			Timestamp:   time.Now(),
			ContentType: "text/plain",
			Body:        []byte(relation),
		})
}

// Consumer 消费消息
func (a *Actor) Consumer(ctx context.Context) {
	// 队列名 持久化 自动删除 exclusive 是否阻塞 额外参数
	_, err := a.channel.QueueDeclare(a.queueName,
		false, false, false, false, nil)

	if err != nil {
		klog.Errorf("cannot declare queue: %v", err)
	}

	//队列名 消费者名 自动应答(自动的进行削峰处理) exclusive 自产自消 阻塞 额外参数
	relations, err := a.channel.Consume(a.queueName, "",
		true, false, false, false, nil)
	if err != nil {
		klog.Errorf("cannot consume queue: %v", err)
	}
	var forever chan struct{}

	// 开启消费消息协程
	go a.RelationActionAdd(ctx, relations)

	<-forever
}
