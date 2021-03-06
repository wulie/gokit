//+build wireinject

package main

import (
	"github.com/google/wire"
)

// wire.go

func InitializeEvent(phrase string) (Event, error) {
	//wire.Build(NewEvent, NewGreeter, NewMessage)
	wire.Build(SuperSet)
	return Event{}, nil
}

func InitializeGreeter(phrase string) (Greeter, error) {
	//wire.Build(NewEvent, NewGreeter, NewMessage)
	wire.Build(SuperSet)
	return Greeter{}, nil
}

//注入方法 不能多也不能少，而且依赖的返回类型不能一样，可以通过重新定义类型解决这个问题👍🏻
