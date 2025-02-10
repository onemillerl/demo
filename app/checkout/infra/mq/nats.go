package mq

import (
	"github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect("nats://192.168.227.128:4222")
	if err != nil {
		panic(err)
	}
}
