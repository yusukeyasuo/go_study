package main

import (
	"fmt"
	"go_study/zoo/animals"
)

func main() {
	fmt.Println(AppName())
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}