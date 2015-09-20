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
	z := make([]string, 5)
	z[0] = "January,5"
	z[1] = "January,15"

	z[2] = "February,5"

	z[3] = "March,5"

	z[4] = "April,25"

	result := commands.SetHolidays(&netNum, &port, &timeOut, &baudRate, z)
	fmt.Println(result)
}
