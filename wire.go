//+build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(string) Event {
   wire.Build(GetMessage, GetGreeter, GetEvent)
   return Event{}
}