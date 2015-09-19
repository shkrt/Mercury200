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
	z := make([]string, 4)
	z[0] = "T1"
	z[1] = "T2"
	z[3] = "T4"
	z[2] = "TSumm"
	result := commands.SetTariffsDisplayOptions(&netNum, &port, &timeOut, &baudRate, z)
	fmt.Println(result)
}
