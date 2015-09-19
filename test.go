package main

import (
	"fmt"
	"mercury200/commands"
	//"time"
)

func main() {
	netNum := "159273"
	port := "COM12"
	timeOut := 5
	baudRate := 9600
	result := commands.GetTariffsDisplayOptions(&netNum, &port, &timeOut, &baudRate)
	fmt.Println(result)
}
