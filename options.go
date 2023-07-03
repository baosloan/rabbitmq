package rabbitmq

type Options struct {
	Host      string //域名或IP
	Port      int    //端口号
	Username  string //用户名
	Password  string //密码
	QueueName string //队列名称
	Exchange  string //交换机
	Key       string //RoutingKey/BindingKey
	Type      Type   //RabbitMQ类型
}

type Type string

const (
	Simple    Type = "simple"     //普通类型
	WorkQueue Type = "work_queue" //工作队列
	PubSub    Type = "pub_sub"    //发布订阅
	Route     Type = "route"      //路由类型
	Topic     Type = "topic"      //话题类型
)

func (opt *Options) init() {
	if opt.Host == "" {
		opt.Host = "127.0.0.1"
	}
	if opt.Port == 0 {
		opt.Port = 5672
	}
	if opt.Type == "" {
		opt.Type = Simple
	}
}
