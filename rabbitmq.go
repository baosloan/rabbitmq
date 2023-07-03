package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	//连接
	conn *amqp.Connection
	//通道
	channel *amqp.Channel
	//队列名称
	queueName string
	//交换机
	exchange string
	//RoutingKey/BindingKey
	key string
	//类型
	rType Type
}

func New(opt *Options) (mq *RabbitMQ, err error) {
	mq = &RabbitMQ{
		queueName: opt.QueueName,
		exchange:  opt.Exchange,
		key:       opt.Key,
		rType:     opt.Type,
	}

	//参数初始化
	opt.init()

	//创建链接
	url := fmt.Sprintf("amqp://%s:%s@%s:%d", opt.Username, opt.Password, opt.Host, opt.Port)
	mq.conn, err = amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	//创建通道
	mq.channel, err = mq.conn.Channel()
	if err != nil {
		return nil, err
	}
	return mq, nil
}

// Close 断开conn和channel
func (r *RabbitMQ) Close() (err error) {
	if err = r.channel.Close(); err != nil {
		return err
	}
	if err = r.conn.Close(); err != nil {
		return err
	}
	return nil
}

type HandleFunc func(msg string) bool

// Publish 生产者客户端接口
func (r *RabbitMQ) Publish(msg string) error {
	switch r.rType {
	case Simple:
		return r.simplePublish(msg)
	case WorkQueue:
	case PubSub:
	case Route:
	case Topic:
	}
	return nil
}

// Consume 消费者客户端接口
func (r *RabbitMQ) Consume(handleFunc HandleFunc) error {
	switch r.rType {
	case Simple:
		return r.simpleConsume(handleFunc)
	case WorkQueue:
	case PubSub:
	case Route:
	case Topic:
	}
	return nil
}
