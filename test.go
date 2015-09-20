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
	z := make([]string, 16)
	z[0] = "January,5"
	z[1] = "January,15"
	z[2] = "February,5"
	z[3] = "March,5"
	z[4] = "April,25"
	z[5] = "May,5"
	z[6] = "June,15"
	z[7] = "July,5"
	z[8] = "August,5"
	z[9] = "September,25"
	z[10] = "October,5"
	z[11] = "November,15"
	z[12] = "December,5"
	z[13] = "May,15"
	z[14] = "April,15"
	z[15] = "January,25"

	result, _ := commands.SetHolidays(&netNum, &port, &timeOut, &baudRate, z)
	fmt.Println(result)
	//fmt.Print(17 % 16)
}
