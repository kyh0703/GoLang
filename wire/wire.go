//go:build wireinject

package main

import (
	"github.com/google/wire"
)

func Initialize() string {
	panic(wire.Build(PSet))
}

var PSet = wire.NewSet(
	provideBar,
	wire.Bind(new(Fooer), new(*MyFooer)),
	provideMyFooer,
)
