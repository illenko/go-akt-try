package main

import (
	"context"

	"github.com/tochemey/goakt/v2/actors"
	"github.com/tochemey/goakt/v2/log"
	"goakt-try/proto"
)

type Driver struct{}

var _ actors.Actor = (*Driver)(nil)

func NewDriver() *Driver {
	return &Driver{}
}

func (d *Driver) PreStart(context.Context) error {
	return nil
}

func (d *Driver) Receive(ctx *actors.ReceiveContext) {
	switch msg := ctx.Message().(type) {
	case *proto.RideRequest:
		log.DefaultLogger.Infof("Driver received ride request to %s", msg.Destination)
		ctx.Tell(ctx.Sender(), &proto.RideAccepted{OrderId: msg.OrderId, DriverId: "driver-1"})
	default:
		ctx.Unhandled()
	}
}

func (d *Driver) PostStop(context.Context) error {
	return nil
}
