package commands

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"mercury200/util"
	"strconv"
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

func PrepareSetterCommand(netNumber *string, code byte, info *[]byte) []byte {
	command := make([]byte, 1)
	command[0] = 0
	command = append(command, util.NetNumToArr(*netNumber)...)
	command = append(command, code)
	command = append(command, *info...)
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

// GET COMMANDS

func GetVersion(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 40)
	val, res := PerformCommand(command, portname, timeout, baud, 13)
	if res == true {
		return fmt.Sprintf("%0x.%0x", val[5], val[6])
	} else {
		return "FAIL"
	}
}

func GetSerial(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 47)
	val, res := PerformCommand(command, portname, timeout, baud, 11)
	if res == true {
		s := fmt.Sprintf("%02x%02x%02x%02x", val[5], val[6], val[7], val[8])
		d, _ := strconv.ParseInt(s, 16, 64)
		return fmt.Sprint(d)
	} else {
		return "FAIL"
	}
}

func GetBatteryVoltage(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 41)
	val, res := PerformCommand(command, portname, timeout, baud, 9)
	if res == true {
		return fmt.Sprintf("%0x.%0x V", val[5], val[6])
	} else {
		return "FAIL"
	}
}

func GetProductionDate(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 102)
	val, res := PerformCommand(command, portname, timeout, baud, 10)
	if res == true {
		return fmt.Sprintf("%02x.%02x.%02x", val[5], val[6], val[7])
	} else {
		return "FAIL"
	}
}

func GetLastTurnOnTime(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 44)
	val, res := PerformCommand(command, portname, timeout, baud, 14)
	if res == true {
		return fmt.Sprintf("%02x.%02x.%02x %02x:%02x:%02x %s", val[9], val[10], val[11], val[6], val[7], val[8], time.Weekday(val[5]))
	} else {
		return "FAIL"
	}
}

func GetLastTurnOffTime(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 43)
	val, res := PerformCommand(command, portname, timeout, baud, 14)
	if res == true {
		return fmt.Sprintf("%02x.%02x.%02x %02x:%02x:%02x %s", val[9], val[10], val[11], val[6], val[7], val[8], time.Weekday(val[5]))
	} else {
		return "FAIL"
	}
}

func GetCurrentTime(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 33)
	val, res := PerformCommand(command, portname, timeout, baud, 14)
	if res == true {
		return fmt.Sprintf("%02x.%02x.%02x %02x:%02x:%02x %s", val[9], val[10], val[11], val[6], val[7], val[8], time.Weekday(val[5]))
	} else {
		return "FAIL"
	}
}

func GetSeasonSwitchFlag(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 36)
	val, res := PerformCommand(command, portname, timeout, baud, 8)
	if res == true {
		if val[5] == 0 {
			return "Switch is disabled"
		} else {
			return "Switch is enabled"
		}
	} else {
		return "FAIL"
	}
}

func GetManualCorrectionAmount(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 37)
	val, res := PerformCommand(command, portname, timeout, baud, 8)
	if res == true {
		return fmt.Sprintf("%02d", val[5])
	} else {
		return "FAIL"
	}
}

//SET COMMANDS

func SetCurrentTime(netNumber *string, portname *string, timeout *int, baud *int, timeToSet time.Time) bool {
	timeP := make([]byte, 7)
	timeP[0] = byte(timeToSet.Weekday())
	timeP[1] = byte(timeToSet.Hour())
	timeP[2] = byte(timeToSet.Minute())
	timeP[3] = byte(timeToSet.Second())
	timeP[4] = byte(timeToSet.Day())
	timeP[5] = byte(timeToSet.Month())
	fmt.Println(timeToSet.Year())
	timeP[6] = byte(timeToSet.Year())
	fmt.Println(timeP)
	//command := PrepareSetterCommand(netNumber, 2, &timeP)
	return false
}
