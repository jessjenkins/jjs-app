package main

import (
	"fmt"

	jj_service "github.com/jessjenkins/jj-service"
	"github.com/jessjenkins/jj-service-app/services"
)

func main() {
	fmt.Println("Starting example app")

	service := jj_service.NewService()

	a := service.AddSubService("a",services.NewUseless("A"))
	b := service.AddSubService("b",services.NewUseless("B"), a)
	c := service.AddSubService("c",services.NewUseless("C"), a)
	service.AddSubService("d", services.NewUseless("D"), b, c)
	service.AddSubService("e", services.NewUseless("E"))

	service.Run()

	fmt.Println("Closing example app")
}