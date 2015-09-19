package commands

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"mercury200/util"
	"strconv"
	"time"
)

type DisplayIntervals struct {
	InactiveTEnergy, ActiveTEnergy, Instants, Additionals int
}

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
		return fmt.Sprintf("%0x.%0x.%0x (%02x.%02x.%02x)", val[5], val[6], val[7], val[8], val[9], val[10])
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

func GetLastOpenedTime(netNumber *string, portname *string, timeout *int, baud *int) (string, error) {
	command := PrepareCommand(netNumber, 97)
	val, res := PerformCommand(command, portname, timeout, baud, 14)
	if res == true {
		if val[5] < 8 {
			return fmt.Sprintf("%02x.%02x.%02x %02x:%02x:%02x %s", val[9], val[10], val[11], val[6], val[7], val[8], time.Weekday(val[5])), nil
		} else {
			return "--:--", nil
		}
	} else {
		return "FAIL", errors.New("No data")
	}
}

func GetLastClosedTime(netNumber *string, portname *string, timeout *int, baud *int) (string, error) {
	command := PrepareCommand(netNumber, 98)
	val, res := PerformCommand(command, portname, timeout, baud, 14)
	if res == true {
		if val[5] < 8 {
			return fmt.Sprintf("%02x.%02x.%02x %02x:%02x:%02x %s", val[9], val[10], val[11], val[6], val[7], val[8], time.Weekday(val[5])), nil
		} else {
			return "--:--", nil
		}
	} else {
		return "FAIL", errors.New("No data")
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

func GetDisplayIntervals(netNumber *string, portname *string, timeout *int, baud *int) DisplayIntervals {
	result := DisplayIntervals{-1, -1, -1, -1}
	p := &result
	command := PrepareCommand(netNumber, 103)
	val, res := PerformCommand(command, portname, timeout, baud, 11)
	if res == true {
		p.InactiveTEnergy = int(val[5])
		p.ActiveTEnergy = int(val[6])
		p.Instants = int(val[7])
		p.Additionals = int(val[8])
		return result
	} else {
		return result
	}
}

func GetTariffsDisplayOptions(netNumber *string, portname *string, timeout *int, baud *int) []string {
	r := make([]string, 5)
	m := make(map[string]string)
	command := PrepareCommand(netNumber, 42)
	val, res := PerformCommand(command, portname, timeout, baud, 8)
	if res == true {
		b := util.SplitEvery(fmt.Sprintf("%0b", val[5]), 1)
		m["T1"] = b[4]
		m["T2"] = b[3]
		m["T3"] = b[2]
		m["T4"] = b[1]
		m["TSumm"] = b[0]
		for k, v := range m {
			if v == "1" {
				r = append(r, k)
			}
		}
		return r
	} else {
		return r
	}
}

//SET COMMANDS

func SetCurrentTime(netNumber *string, portname *string, timeout *int, baud *int, timeToSet time.Time) bool {
	timeP := make([]byte, 7)
	timeP[0] = byte(timeToSet.Weekday())

	h, _ := strconv.ParseInt(strconv.Itoa(timeToSet.Hour()), 16, 64)
	timeP[1] = byte(h)

	h, _ = strconv.ParseInt(strconv.Itoa(timeToSet.Minute()), 16, 64)
	timeP[2] = byte(h)

	h, _ = strconv.ParseInt(strconv.Itoa(timeToSet.Second()), 16, 64)
	timeP[3] = byte(h)

	h, _ = strconv.ParseInt(strconv.Itoa(timeToSet.Day()), 16, 64)
	timeP[4] = byte(h)

	h, _ = strconv.ParseInt(fmt.Sprintf("%d", timeToSet.Month()), 16, 64)
	timeP[5] = byte(h)

	h, _ = strconv.ParseInt(strconv.Itoa(timeToSet.Year())[2:4], 16, 64)
	timeP[6] = byte(h)

	command := PrepareSetterCommand(netNumber, 2, &timeP)
	_, res := PerformCommand(command, portname, timeout, baud, 7)
	if res == true {
		return true
	} else {
		return false
	}
}

func SetSeasonSwitchFlag(netNumber *string, portname *string, timeout *int, baud *int, flag bool) bool {
	tail := make([]byte, 1)

	if flag == true {
		tail[0] = 255
	} else {
		tail[0] = 0
	}

	command := PrepareSetterCommand(netNumber, 5, &tail)
	_, res := PerformCommand(command, portname, timeout, baud, 7)
	if res == true {
		return true
	} else {
		return false
	}
}

func SetManualCorrectionAmount(netNumber *string, portname *string, timeout *int, baud *int, amount uint) (bool, error) {
	tail := make([]byte, 1)

	if amount <= 89 {
		tail[0] = byte(amount)
	} else {
		return false, errors.New("Amount must be between 0 and 89")
	}

	command := PrepareSetterCommand(netNumber, 6, &tail)
	_, res := PerformCommand(command, portname, timeout, baud, 7)
	if res == true {
		return true, nil
	} else {
		return false, nil
	}
}

func SetTariffsDisplayOptions(netNumber *string, portname *string, timeout *int, baud *int, lst []string) bool {
	var s bytes.Buffer
	m := make(map[string]bool)

	for _, v := range lst {
		m[v] = true
	}

	s.WriteString("000")

	if m["TSumm"] {
		s.WriteString("1")
		fmt.Println(s[2])
	} else {
		s.WriteString("0")

	}

	if m["T4"] {
		s.WriteString("1")
	} else {
		s.WriteString("0")
	}

	if m["T3"] {
		s.WriteString("1")
	} else {
		s.WriteString("0")
	}

	if m["T2"] {
		s.WriteString("1")
	} else {
		s.WriteString("0")
	}

	if m["T1"] {
		s.WriteString("1")
	} else {
		s.WriteString("0")
	}

	e, _ := strconv.ParseInt(s.String(), 2, 64)
	tail := make([]byte, 1)
	tail[0] = byte(e)
	command := PrepareSetterCommand(netNumber, 9, &tail)
	_, res := PerformCommand(command, portname, timeout, baud, 7)
	if res == true {
		return true
	} else {
		return false
	}
}
