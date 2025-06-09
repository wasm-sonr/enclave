package runtime

import (
	"log/slog"

	"github.com/asynkron/protoactor-go/actor"
)

type (
	hello struct{ Who string }
)

type EnclaveActor struct {
	host EnclaveHost
}

func (a *EnclaveActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		context.Logger().Info("Hello ", slog.String("who", msg.Who))
	case *actor.Started:
		context.Respond(a.host)
	case *actor.Stopping:
		context.Respond(true)
	case *actor.Stopped:
		context.Respond(true)
	}
}
