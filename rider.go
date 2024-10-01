package main

import (
	"context"

	"github.com/tochemey/goakt/v2/actors"
	"github.com/tochemey/goakt/v2/log"
	"goakt-try/proto"
)

type Rider struct{}

var _ actors.Actor = (*Rider)(nil)

func NewRider() *Rider {
	return &Rider{}
}

func (r *Rider) PreStart(context.Context) error {
	return nil
}

func (r *Rider) Receive(ctx *actors.ReceiveContext) {
	switch msg := ctx.Message().(type) {
	case *proto.RideAccepted:
		log.DefaultLogger.Infof("Ride accepted by driver %s", msg.DriverId)
	default:
		ctx.Unhandled()
	}
}

func (r *Rider) PostStop(context.Context) error {
	return nil
}
