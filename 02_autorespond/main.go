package main

import (
	"fmt"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

type myAutoResponder struct {
	name string
}

func (m *myAutoResponder) GetAutoResponse(context actor.Context) interface{} {
	return &myAutoResponder{
		name: m.name + ": ``" + context.Self().Id + "``",
	}
}

type myAutoRespoce struct {
	name string
}

func main() {
	system := actor.NewActorSystem()
	props := actor.PropsFromFunc(func(ctx actor.Context) {})
	pid := system.Root.Spawn(props)

	res, _ := system.Root.RequestFuture(pid, &myAutoResponder{name: "hello"}, 10*time.Second).Result()

	fmt.Printf("%v", res)
}
