package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tochemey/goakt/v2/actors"
	"github.com/tochemey/goakt/v2/log"
	"goakt-try/proto"
)

func main() {
	ctx := context.Background()

	logger := log.DefaultLogger

	actorSystem, _ := actors.NewActorSystem("TaxiOrderingSystem",
		actors.WithLogger(logger),
		actors.WithPassivationDisabled(),
		actors.WithActorInitMaxRetries(3))

	_ = actorSystem.Start(ctx)

	time.Sleep(1 * time.Second)

	rider := NewRider()
	driver := NewDriver()
	riderActor, _ := actorSystem.Spawn(ctx, "Rider", rider)
	driverActor, _ := actorSystem.Spawn(ctx, "Driver", driver)

	time.Sleep(1 * time.Second)

	if err := riderActor.Tell(ctx, driverActor, &proto.RideRequest{OrderId: "order-1", RiderId: "rider-1", Destination: "Central Park"}); err != nil {
		panic(err)
	}

	done := make(chan struct{})
	go func() {
		for await := time.After(1 * time.Minute); ; {
			select {
			case <-await:
				done <- struct{}{}
				return
			}
		}
	}()

	<-done

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	_ = actorSystem.Stop(ctx)
	os.Exit(0)
}
