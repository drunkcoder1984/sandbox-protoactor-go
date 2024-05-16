package main

import (
	"log/slog"
	"time"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

type (
	hello      struct{ Who string }
	helloActor struct{}
)

func (state *helloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		context.Logger().Info("Starting, initialize actor here")
	case *actor.Stopping:
		context.Logger().Info("Stopping, actor is about to shut down")
	case *actor.Stopped:
		context.Logger().Info("Stopped, actor and its children are stopped")
	case *actor.Restarting:
		context.Logger().Info("Restarting, actor is about to restart")
	case *hello:
		context.Logger().Info("Hello", slog.String("Who", msg.Who))
	}
}

func main() {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(func() actor.Actor { return &helloActor{} })
	pid := system.Root.Spawn(props)
	system.Root.Send(pid, &hello{Who: "Roger"})

	time.Sleep(1 * time.Second)
	system.Root.Stop(pid)

	_, _ = console.ReadLine()
}
