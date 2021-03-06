package main

import (
	"errors"
	"fmt"
	"github.com/google/wire"
	"os"
	"time"
)

func main() {
	//message := NewMessage()
	//greeter := NewGreeter(message)
	//event := NewEvent(greeter)
	//
	//event.Start()

	//e := InitializeEvent()
	//
	//e.Start()

	e, err := InitializeEvent("helloxxxx")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	println(e.Greeter.Greet())
	e.Start()
}

type Message string

func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Greeter struct {
	Message Message // <- adding a Message field
	Grumpy  bool
	phrase  string
}

func NewGreeter(m Message, phrase string) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	fmt.Println("greeter phrase is ", phrase)
	return Greeter{Message: m, Grumpy: grumpy}
}
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

var SuperSet = wire.NewSet(NewMessage, NewEvent, NewGreeter)
