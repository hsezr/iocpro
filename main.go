package main

import (
	"fmt"
	"iocpro/injector"
	"iocpro/services"
)

func main() {
	injector.BeanFactory.Set(services.NewOrderService())
	
	userService := services.NewUserService()
	injector.BeanFactory.Apply(userService)

	fmt.Println(userService.Order)
}