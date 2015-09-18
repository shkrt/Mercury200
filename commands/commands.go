package commands

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"mercury200/util"
	"time"
)

func PrepareCommand(netNumber *string, code byte) []byte {
	command := make([]byte, 1)
	command[0] = 0
	command = append(command, util.NetNumToArr(*netNumber)...)
	command = append(command, code)
	var crc = util.GetCrcBytes(command)
	command = append(command, crc...)
	return command
}

func PerformCommand(command []byte, portname *string, timeout *int, baud *int, respLen int) ([]byte, bool) {

	c := &serial.Config{Name: *portname, Baud: *baud, ReadTimeout: time.Second * time.Duration(*timeout)}

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

	if util.CheckCrc(buf[:n], respLen) {
		return buf[:n], true
	} else {
		return buf, false
	}

}

//COUNTER COMMANDS

func GetVersion(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 40)
	val, res := PerformCommand(command, portname, timeout, baud, 13)
	if res == true {
		return fmt.Sprintf("%0x.%0x", val[5], val[6])
	} else {
		return "FAIL"
	}
}
