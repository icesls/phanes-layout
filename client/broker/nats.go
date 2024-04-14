package broker

import (
	"github.com/go-micro/plugins/v4/broker/nats"
	"go-micro.dev/v4/broker"
	"phanes/config"
	"phanes/utils"
)

var defaultNatsAddress = "nats://127.0.0.1:1222"

func InitNats() broker.Broker {
	if config.Conf.Broker.Addr != "" {
		defaultNatsAddress = config.Conf.Broker.Addr
	}
	b := nats.NewBroker(broker.Addrs(defaultNatsAddress))
	utils.Throw(b.Connect())
	return b
}
