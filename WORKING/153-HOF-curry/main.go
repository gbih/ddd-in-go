package main

import "fmt"

// function with two parameters
func sayGreetingTwoParam(greeting, name string) string {
	return (fmt.Sprintf("%s %s", greeting, name))
}

// In Go, we cannot just pass in one parameter to create a function:
// var sayHello = sayGreeting("Hello")

// In Go, we need to specifically create a curried function
func sayGreeting(greeting string) func(string) string {
	return func(name string) string {
		return (fmt.Sprintf("%s %s", greeting, name))
	}
}

// This approach of “baking in” parameters is called
// partial application and is a very important functional pattern.
// We can use it to implement pseudo-dependency injection FP-style.

//------------------------

func main() {
	fmt.Println("------------------------------")
	fmt.Println(sayGreetingTwoParam("Hello", "George"))

	//	sayGreeting("Hello") is our partial application
	var sayHello = sayGreeting("Hello")
	fmt.Printf("partial application sayHello: %#v\n", sayHello)
	fmt.Println(sayHello("Alex"))

	var sayGoodbye = sayGreeting("Goodbye")
	fmt.Printf("partial application sayGoodbye: %#v\n", sayGoodbye)
	fmt.Println(sayGoodbye("Alex"))

}
