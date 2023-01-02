
# Wire_package

What is dependency injection?
Dependency injection is a software engineering technique where an object or struct receives its dependencies at compile time. Wikipedia defines dependency injection as such:

Dependency injection is a technique in which an object receives other objects that it depends on, called dependencies. Typically, the receiving object is called a client and the passed-in (‘injected’) object is called a service.

To get a better view of this, let’s analyze an example. Take a look at the following code:

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

func GetMessage() Message {
   return Message("Hello world!")
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
   message := GetMessage()
   greeter := GetGreeter(message)
   event := GetEvent(greeter)

   event.Start()
}
If you take a look at the code above, we have a message, a greeter, and an event. There is also a GetMessage function that returns a message; a GetGreeter function that takes in a message and returns a greeter; and a GetEvent function that accepts a greeter and returns an event. The event also has a method called Start that prints out the message.

If you take a look at our main method, we first create a message, then we pass in the message as a dependency to the greeter and finally pass that to the event. Run the code by running the command go run . in the terminal.

Go Dependency Injection With Wire
As you can see, it prints “Hello, world!” to the console. This is a very shallow dependency graph, but you can already see the complexity that comes with this when implementing this in a large codebase. That’s where dependency injection tools like Wire come in.
Wire is a code dependency tool that operates without runtime state or reflection.
Code written to be used with Wire is useful even for handwritten initialization.
Wire can generate source code at compile time as well as implement dependency injection. 

According to the official documentation, “In Wire, dependencies between components are represented as function parameters, 
encouraging explicit initialization instead of global variables.”

To use Wire, first, you need to initialize Go modules in your current working directory. Run the command go mod init go-wire to do this.
![image](https://user-images.githubusercontent.com/114211073/210212451-e44b39e4-3cff-4351-91a3-e826ee8dc0a7.png)
Now, run the command go get github.com/google/wire/cmd/wire to install.
![image](https://user-images.githubusercontent.com/114211073/210212559-ff603989-ccfd-448e-861d-dc8328d63eea.png)
Now, let’s refactor our code to use Wire as a dependency injection tool. Create a file called wire.go and add the following code:

package main

import "github.com/google/wire"

func InitializeEvent() Event {
   wire.Build(GetMessage, GetGreeter, GetEvent)
   return Event{}
}

Now, let’s refactor our code to use Wire as a dependency injection tool. Create a file called wire.py and add the following code:

package main

import "github.com/google/wire"

func InitializeEvent() Event {
   wire.Build(GetMessage, GetGreeter, GetEvent)
   return Event{}
}
First of all, we import Wire, then we create a function called InitializeEvent. This function returns an event that we will use in our main method. In the InitializeEvent function, we make a call to Wire. Then we build and pass in all our dependencies. Note we can pass in these dependencies in any order.

Then we return an empty event. Don’t worry, Wire will take over here!

Now, change your main method to this:

func main() {
   event := InitializeEvent()
   event.Start()
}
Notice how we have successfully cut down the code in our main method to just two lines.
Run the command go run github.com/google/wire/cmd/wire to generate our dependencies with Wire.

![image](https://user-images.githubusercontent.com/114211073/210212802-d59ef0d8-c014-424e-8498-d7701f6957a0.png)
Now you will see that Wire has generated a file called wire_gen.

![image](https://user-images.githubusercontent.com/114211073/210212864-0bfca674-6c88-4bb8-b2bb-3a7932e9a313.png)
If you should run the code again, you will get an error.
![image](https://user-images.githubusercontent.com/114211073/210212893-23de225e-a6c1-4877-af5c-353430d6c5f0.png)
This is because our InitializeEvent function has now been redeclared in the wire_gen file. Add
//+build wireinject to the beginning of your wire.go file to tell Go to ignore it when building. Make sure to add a new line after that or this will not work.
![image](https://user-images.githubusercontent.com/114211073/210212920-9bc5ae5b-953b-4020-a553-8dc420ef3af3.png)
If you run go run . again, you should still see the same “Hello, world!” output.
![image](https://user-images.githubusercontent.com/114211073/210212944-6e7b4971-134b-4522-b7ef-3efea33cb976.png)


Working with arguments
What if you wanted to dynamically pass in a message as an argument? Let’s take a look at how we can do this. Modify the GetMessage function to this:
func GetMessage(text string) Message {
   return Message(text)
}
Now we have to pass in a text int the paranthesis of initialize event in the main.go file to display. Let’s try running this and see the output.
![image](https://user-images.githubusercontent.com/114211073/210213055-5b1af02d-7871-47e9-b97c-3fadaed536df.png)

As you can see, Wire recognizes that we have to pass in an argument to the GetMessage function. Let’s resolve this error. Modify your InitializeEvent function in your wire.go file:

func InitializeEvent(string) Event {
   wire.Build(GetMessage, GetGreeter, GetEvent)
   return Event{}
}
Now we are telling Wire that we expect a string argument.

Run go run github.com/google/wire/cmd/wire again. If you take a look at our wire_gen.go file, you will see that Wire has refactored the code to accept this value.

