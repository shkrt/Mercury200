package main

import (
	"fmt"
	"mercury200/commands"
	"time"
)

func main() {
	netNum := "159273"
	port := "COM12"
	timeOut := 5
	baudRate := 9600
	result := commands.SetCurrentTime(&netNum, &port, &timeOut, &baudRate, time.Now())
	fmt.Println(result)
}
