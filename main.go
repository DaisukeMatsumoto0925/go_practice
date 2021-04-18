package main

import "fmt"

type (
	Message string
	Greeter struct {
		Message Message
	}

	Event struct {
		Greeter Greeter
	}
)

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	message := NewMessage()
	greeter := NewGreter(message)
	event := NewEvent(greeter)

	event.Start()
}
