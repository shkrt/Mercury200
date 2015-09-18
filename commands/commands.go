package commands

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"mercury200/util"
	"time"
)

func PrepareCommand(netNumber string, code byte) []byte {
	command := make([]byte, 1)
	command[0] = 0
	command = append(command, util.NetNumToArr(netNumber)...)
	command = append(command, code)
	var crc = util.GetCrcBytes(command)
	command = append(command, crc...)
	fmt.Println(command)
	return command
}

func PerformCommand(command []byte, netNumber string, portname string, timeout int, baud int, respLen int) bool {

	c := &serial.Config{Name: portname, Baud: baud, ReadTimeout: time.Second * time.Duration(timeout)}

	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write(command)

	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, respLen)
	n, err = s.Read(buf)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%q", buf[:n])
	if util.CheckCrc(buf[:n], respLen) {
		return true
	} else {
		return false
	}

}

//COUNTER COMMANDS

func GetVersion(netNumber string, portname string, timeout int, baud int) {
	command := PrepareCommand(netNumber, 40)
	PerformCommand(command, netNumber, portname, timeout, baud, 13)
}
