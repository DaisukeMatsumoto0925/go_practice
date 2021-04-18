package main

type (
	Message string
	Greeter struct {
		Message Message
	}

	Event struct {
	}
)

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreter(m Message) Greeter {
	return Greeter{Message: m}
}
