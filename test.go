package main

import (
	"fmt"
	"mercury200/commands"
	"time"
)

func main() {
	netNum := "266608"
	port := "COM5"
	timeOut := 5
	baudRate := 9600
	result := commands.SetCurrentTime(&netNum, &port, &timeOut, &baudRate, time.Now())
	fmt.Println(result)
}
