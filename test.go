package main

import (
	"fmt"
	"mercury200/commands"
	//"mercury200/types"
)

func main() {
	netNum := "159273"
	port := "COM12"
	timeOut := 5
	baudRate := 9600

	result, _ := commands.GetHolidays(&netNum, &port, &timeOut, &baudRate)
	fmt.Println(result)
	//fmt.Print(17 % 16)
}
