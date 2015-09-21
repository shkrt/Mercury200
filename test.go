package main

import (
	"fmt"
	"mercury200/commands"
)

func main() {
	netNum := "266608"
	port := "COM5"
	timeOut := 5
	baudRate := 9600

	result, _ := commands.GetEnergyAtMonthStart(&netNum, &port, &timeOut, &baudRate, 11)
	fmt.Println(result)

	fmt.Printf("Тариф 1: %s кВт\n", result.T1)
	fmt.Printf("Тариф 2: %s кВт", result.T2)
}
