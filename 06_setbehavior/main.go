package main

import (
	"fmt"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

type (
	Hello            struct{ Who string }
	SetBehaviorActor struct {
		behavior actor.Behavior
	}
)

func (state *SetBehaviorActor) Receive(context actor.Context) {
	state.behavior.Receive(context)
}

func (state *SetBehaviorActor) One(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
		// 次のメッセージを受け取る時に実行されるメソッドをstate.Otherに変更
		state.behavior.Become(state.Other)
	}
}

func (state *SetBehaviorActor) Other(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("%v, ey we are now handling messages in another behavior", msg.Who)
	}
}

func NewSetBehaviorActor() actor.Actor {
	act := &SetBehaviorActor{
		behavior: actor.NewBehavior(),
	}
	// デフォルトのReceiveメソッドをOneに設定
	act.behavior.Become(act.One)
	return act
}

func main() {
	system := actor.NewActorSystem()
	rootContext := system.Root
	props := actor.PropsFromProducer(NewSetBehaviorActor)
	pid := rootContext.Spawn(props)
	rootContext.Send(pid, Hello{Who: "Roger"}) // SetBehaviorActor.Oneが実行される
	rootContext.Send(pid, Hello{Who: "Hoge"})  // SetBehaviorActor.Otherが実行される
	_, _ = console.ReadLine()
}
