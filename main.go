package main

import (
	"fmt"

	"github.com/jessjenkins/jjs"
	"github.com/jessjenkins/jjs-app/services"
)

func main() {
	fmt.Println("Starting example app")

	runner := jjs.NewRunner()

	a := runner.AddService("a", services.NewUseless("A"))
	b := runner.AddService("b", services.NewUseless("B"), a)
	c := runner.AddService("c", services.NewUseless("C"), a)
	runner.AddService("d", services.NewUseless("D"), b, c)
	runner.AddService("e", services.NewUseless("E"))

	runner.Run()

	fmt.Println("Closing example app")
}
