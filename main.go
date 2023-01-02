package main

import (
   "fmt"
)

type Message string
type Greeter struct {
   Message Message
}
type Event struct {
   Greeter Greeter
}


func GetMessage(text string) Message {
   return Message(text)
}
func GetGreeter(m Message) Greeter {
   return Greeter{Message: m}
}
func (g Greeter) Greet() Message {
   return g.Message
}
func GetEvent(g Greeter) Event {
   return Event{Greeter: g}
}
func (e Event) Start() {
   msg := e.Greeter.Greet()
   fmt.Println(msg)
}
func main() {
	event := InitializeEvent("hello world")
	event.Start()
 }