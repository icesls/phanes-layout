package broker

import (
	"github.com/go-micro/plugins/v4/broker/rabbitmq"
	"go-micro.dev/v4/broker"
	"phanes/config"
	"phanes/utils"
)

var defaultRabbitMQAddress = "amqp://guest:guest@localhost:5672"

func InitRabbit() broker.Broker {
	if !config.Conf.Broker.Enabled {
		return nil
	}
	if config.Conf.Broker.Addr != "" {
		defaultRabbitMQAddress = config.Conf.Broker.Addr
	}
	b := rabbitmq.NewBroker(broker.Addrs(defaultRabbitMQAddress))
	utils.Throw(b.Connect())
	return b
}
