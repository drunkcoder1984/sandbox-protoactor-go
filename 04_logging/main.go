package main

import (
	"log/slog"
	"os"
	"time"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/lmittmann/tint"
	slogzap "github.com/samber/slog-zap/v2"
	"go.uber.org/zap"
)

type (
	hello      struct{ Who string }
	helloActor struct{}
)

func (state *helloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		context.Logger().Info("Hello ", slog.String("Who", msg.Who))
	}
}

func jsonLogging(system *actor.ActorSystem) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil)).
		With("lib", "Prot.Actor").
		With("system", system.ID)
}

func consoleLogging(system *actor.ActorSystem) *slog.Logger {
	return slog.Default().
		With("lib", "Prot.Actor").
		With("system", system.ID)
}

func coloredConsoleLogging(system *actor.ActorSystem) *slog.Logger {
	return slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.RFC3339,
		AddSource:  true,
	})).With("lib", "Prot.Actor").With("system", system.ID)
}

func zapAdapterLogging(system *actor.ActorSystem) *slog.Logger {
	zapLogger, _ := zap.NewProduction()

	logger := slog.New(slogzap.Option{Level: slog.LevelDebug, Logger: zapLogger}.NewZapHandler())
	return logger.
		With("lib", "Prot.Actor").
		With("system", system.ID)
}

func main() {
	run(jsonLogging, "jsonLogging")
	run(consoleLogging, "consoleLogging")
	run(consoleLogging, "coloredConsoleLogging")
	run(zapAdapterLogging, "zapAdapterLogging")

	_, _ = console.ReadLine()
}

func run(factory func(system *actor.ActorSystem) *slog.Logger, who string) {
	system := actor.NewActorSystem(actor.WithLoggerFactory(factory))
	props := actor.PropsFromProducer(func() actor.Actor { return &helloActor{} })

	pid := system.Root.Spawn(props)
	system.Root.Send(pid, &hello{Who: who})
}
