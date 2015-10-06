package commands

import (
	"errors"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"mercury200/types"
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
		s.Close()
		return buf[:n], true
	} else {
		s.Close()
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
		return ""
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
		return ""
	}
}

func GetBatteryVoltage(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 41)
	val, res := PerformCommand(command, portname, timeout, baud, 9)
	if res == true {
		return fmt.Sprintf("%0x.%0x V", val[5], val[6])
	} else {
		return ""
	}
}

func GetProductionDate(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 102)
	val, res := PerformCommand(command, portname, timeout, baud, 10)
	if res == true {
		return fmt.Sprintf("%02x.%02x.%02x", val[5], val[6], val[7])
	} else {
		return ""
	}
}

func GetLastTurnOnTime(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 44)
	val, res := PerformCommand(command, portname, timeout, baud, 14)
	if res == true {
		return fmt.Sprintf("%02x.%02x.%02x %02x:%02x:%02x %s",
			val[9], val[10], val[11], val[6], val[7], val[8], time.Weekday(val[5]))
	} else {
		return ""
	}
}

func GetLastTurnOffTime(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 43)
	val, res := PerformCommand(command, portname, timeout, baud, 14)
	if res == true {
		return fmt.Sprintf("%02x.%02x.%02x %02x:%02x:%02x %s",
			val[9], val[10], val[11], val[6], val[7], val[8], time.Weekday(val[5]))
	} else {
		return ""
	}
}

func GetCurrentTime(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 33)
	val, res := PerformCommand(command, portname, timeout, baud, 14)
	if res == true {
		return fmt.Sprintf("%02x.%02x.%02x %02x:%02x:%02x %s",
			val[9], val[10], val[11], val[6], val[7], val[8], time.Weekday(val[5]))
	} else {
		return ""
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
		return ""
	}
}

func GetLastOpenedTime(netNumber *string, portname *string, timeout *int, baud *int) (string, error) {
	command := PrepareCommand(netNumber, 97)
	val, res := PerformCommand(command, portname, timeout, baud, 14)
	if res == true {
		if val[5] < 8 {
			return fmt.Sprintf("%02x.%02x.%02x %02x:%02x:%02x %s",
				val[9], val[10], val[11], val[6], val[7], val[8], time.Weekday(val[5])), nil
		} else {
			return "--:--", nil
		}
	} else {
		return "", errors.New("No data")
	}
}

func GetLastClosedTime(netNumber *string, portname *string, timeout *int, baud *int) (string, error) {
	command := PrepareCommand(netNumber, 98)
	val, res := PerformCommand(command, portname, timeout, baud, 14)
	if res == true {
		if val[5] < 8 {
			return fmt.Sprintf("%02x.%02x.%02x %02x:%02x:%02x %s",
				val[9], val[10], val[11], val[6], val[7], val[8], time.Weekday(val[5])), nil
		} else {
			return "--:--", nil
		}
	} else {
		return "", errors.New("No data")
	}
}

func GetManualCorrectionAmount(netNumber *string, portname *string, timeout *int, baud *int) string {
	command := PrepareCommand(netNumber, 37)
	val, res := PerformCommand(command, portname, timeout, baud, 8)
	if res == true {
		return fmt.Sprintf("%02d", val[5])
	} else {
		return ""
	}
}

func GetDisplayIntervals(netNumber *string, portname *string, timeout *int, baud *int) types.DisplayIntervals {
	result := types.DisplayIntervals{-1, -1, -1, -1}
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

func GetTariffsDisplayOptions(netNumber *string, portname *string, timeout *int, baud *int) *types.TariffsDisplayOptions {
	c := types.TariffsDisplayOptions{}
	p := &c

	command := PrepareCommand(netNumber, 42)
	val, res := PerformCommand(command, portname, timeout, baud, 8)
	if res == true {
		b := util.SplitEvery(fmt.Sprintf("%08b", val[5]), 1)
		p.Date = b[0]
		p.Time = b[1]
		p.Power = b[2]
		p.TSumm = b[3]
		p.T4 = b[4]
		p.T3 = b[5]
		p.T2 = b[6]
		p.T1 = b[7]

		return p
	} else {
		return p
	}
}

func GetPowerLimit(netNumber *string, portname *string, timeout *int, baud *int) int {
	command := PrepareCommand(netNumber, 34)
	val, res := PerformCommand(command, portname, timeout, baud, 9)
	if res == true {
		k, _ := strconv.Atoi(fmt.Sprintf("%0x%0x", val[5], val[6]))
		return k * 10
	}
	return -1
}

func GetEnergyLimit(netNumber *string, portname *string, timeout *int, baud *int) int {
	command := PrepareCommand(netNumber, 35)
	val, res := PerformCommand(command, portname, timeout, baud, 9)
	if res == true {
		k, _ := strconv.Atoi(fmt.Sprintf("%0x%0x", val[5], val[6]))
		return k
	}
	return -1
}

func GetImpOutputOptions(netNumber *string, portname *string, timeout *int, baud *int) string {
	options := map[byte]string{
		0: "5000 imp/h",
		1: "10000 imp/h",
		2: "Quartz frequency",
		3: "Load control",
	}

	command := PrepareCommand(netNumber, 45)
	val, res := PerformCommand(command, portname, timeout, baud, 8)
	if res == true {
		return options[val[5]]
	}
	return ""
}

func GetTariffsCount(netNumber *string, portname *string, timeout *int, baud *int) int {
	command := PrepareCommand(netNumber, 46)
	val, res := PerformCommand(command, portname, timeout, baud, 8)
	if res == true {
		return int(val[5])
	}
	return -1
}

func GetHolidays(netNumber *string, portname *string, timeout *int, baud *int) ([]string, error) {
	result := make([]string, 0)
	tail := make([]byte, 1)
	for i := 0; i < 2; i++ {
		tail[0] = byte(i)
		command := PrepareSetterCommand(netNumber, 48, &tail)
		val, res := PerformCommand(command, portname, timeout, baud, 23)
		if res == true {
			for j := 5; j < 21; j += 2 {
				if val[j] != 255 {
					m, _ := strconv.Atoi(fmt.Sprintf("%x", val[j+1]))
					h := fmt.Sprintf("%s,%0x", time.Month(m), val[j])
					result = append(result, h)
				} else {
					return result, nil
				}
			}
		} else {
			return result, errors.New("unable to fetch holidays data")
		}
	}

	return result, nil
}

func GetEnergyFromReset(netNumber *string, portname *string, timeout *int, baud *int) *types.Energy {
	energy := types.Energy{}
	p := &energy
	command := PrepareCommand(netNumber, 39)
	val, res := PerformCommand(command, portname, timeout, baud, 23)
	if res == true {
		p.T1 = fmt.Sprintf("%x%x%x.%x", val[5], val[6], val[7], val[8])
		p.T2 = fmt.Sprintf("%x%x%x.%x", val[9], val[10], val[11], val[12])
		p.T3 = fmt.Sprintf("%x%x%x.%x", val[13], val[14], val[15], val[16])
		p.T4 = fmt.Sprintf("%x%x%x.%x", val[17], val[18], val[19], val[20])
	}
	return p
}

func GetEnergyAtMonthStart(netNumber *string, portname *string, timeout *int, baud *int, month int) (*types.Energy, error) {
	energy := types.Energy{}
	p := &energy

	if month < 1 || month > 12 {
		return p, errors.New("month should be between 0 and 12")
	}

	tail := make([]byte, 1)
	tail[0] = byte(month - 1)
	command := PrepareSetterCommand(netNumber, 50, &tail)
	val, res := PerformCommand(command, portname, timeout, baud, 23)

	if res == true {
		p.T1 = fmt.Sprintf("%x%x%x.%x", val[5], val[6], val[7], val[8])
		p.T2 = fmt.Sprintf("%x%x%x.%x", val[9], val[10], val[11], val[12])
		p.T3 = fmt.Sprintf("%x%x%x.%x", val[13], val[14], val[15], val[16])
		p.T4 = fmt.Sprintf("%x%x%x.%x", val[17], val[18], val[19], val[20])
	}else{
		return p, errors.New("CRC Check error")
	}


	return p, nil
}

func GetInstants(netNumber *string, portname *string, timeout *int, baud *int) *types.Instants {
	values := types.Instants{}
	p := &values
	command := PrepareCommand(netNumber, 99)
	val, res := PerformCommand(command, portname, timeout, baud, 14)
	if res == true {
		v, _ := strconv.ParseFloat(fmt.Sprintf("%x%x", val[5], val[6]), 32)
		p.U = fmt.Sprint(v / 10)
		v, _ = strconv.ParseFloat(fmt.Sprintf("%x%x", val[7], val[8]), 32)
		p.I = fmt.Sprint(v / 10)
		v, _ = strconv.ParseFloat(fmt.Sprintf("%x%x%x", val[9], val[10], val[11]), 32)
		p.P = fmt.Sprint(v)
	}
	return p

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

func SetTariffsDisplayOptions(netNumber *string, portname *string, timeout *int, baud *int, opt *types.TariffsDisplayOptions) bool {
	s := []byte("00000000")

	if opt.Date == "1" {
		s[0] = 49
	}

	if opt.Time == "1" {
		s[1] = 49
	}

	if opt.Power == "1" {
		s[2] = 49
	}

	if opt.TSumm == "1" {
		s[3] = 49
	}

	if opt.T4 == "1" {
		s[4] = 49
	}

	if opt.T3 == "1" {
		s[5] = 49
	}

	if opt.T2 == "1" {
		s[6] = 49
	}

	if opt.T1 == "1" {
		s[7] = 49
	}

	e, _ := strconv.ParseInt(string(s), 2, 64)
	tail := make([]byte, 1)
	tail[0] = byte(e)
	command := PrepareSetterCommand(netNumber, 9, &tail)
	_, res := PerformCommand(command, portname, timeout, baud, 7)
	if res == true {
		return true
	}
	return false
}

func SetDisplayIntervals(netNumber *string, portname *string, timeout *int, baud *int, intervals *types.DisplayIntervals) bool {
	tail := make([]byte, 4)
	tail[0] = byte(intervals.InactiveTEnergy)
	tail[1] = byte(intervals.ActiveTEnergy)
	tail[2] = byte(intervals.Instants)
	tail[3] = byte(intervals.Additionals)

	command := PrepareSetterCommand(netNumber, 13, &tail)
	_, res := PerformCommand(command, portname, timeout, baud, 7)
	if res == true {
		return true
	}
	return false
}

func SetPowerLimit(netNumber *string, portname *string, timeout *int, baud *int, limit int) (bool, error) {
	tail := make([]byte, 2)

	if limit <= 99999 && limit > 0 {
		limit = limit / 10
		h := util.SplitEvery(fmt.Sprintf("%04d", limit), 2)
		var l int64

		for i, v := range h {
			l, _ = strconv.ParseInt(v, 16, 64)
			tail[i] = byte(l)
		}

	} else {
		return false, errors.New("Amount must be between 0 and 99999")
	}

	command := PrepareSetterCommand(netNumber, 3, &tail)
	_, res := PerformCommand(command, portname, timeout, baud, 7)
	if res == true {
		return true, nil
	}
	return false, nil
}

func SetEnergyLimit(netNumber *string, portname *string, timeout *int, baud *int, limit int) (bool, error) {
	tail := make([]byte, 2)

	if limit <= 9999 && limit > 0 {
		h := util.SplitEvery(fmt.Sprintf("%04d", limit), 2)
		var l int64

		for i, v := range h {
			l, _ = strconv.ParseInt(v, 16, 64)
			tail[i] = byte(l)
		}

	} else {
		return false, errors.New("Amount must be between 0 and 9999")
	}

	command := PrepareSetterCommand(netNumber, 4, &tail)
	_, res := PerformCommand(command, portname, timeout, baud, 7)
	if res == true {
		return true, nil
	}
	return false, nil
}

func SetImpOutputOptions(netNumber *string, portname *string, timeout *int, baud *int, option int) bool {
	tail := make([]byte, 1)

	tail[0] = byte(option)

	command := PrepareSetterCommand(netNumber, 7, &tail)
	_, res := PerformCommand(command, portname, timeout, baud, 7)
	if res == true {
		return true
	}
	return false
}

func SetTariffsCount(netNumber *string, portname *string, timeout *int, baud *int, option int) (bool, error) {
	tail := make([]byte, 1)

	if option <= 4 && option > 0 {
		tail[0] = byte(option)

	} else {
		return false, errors.New("option must be between 1 and 4")
	}

	command := PrepareSetterCommand(netNumber, 10, &tail)
	_, res := PerformCommand(command, portname, timeout, baud, 7)
	if res == true {
		return true, nil
	}

	return false, nil
}

func SetHolidays(netNumber *string, portname *string, timeout *int, baud *int, holidays []string) (bool, error) {
	tail := make([]byte, 0)
	const shortForm = "January,2"
	if len(holidays) > 16 {
		return false, errors.New("only 16 holidays supported")
	}

	for _, v := range holidays {
		if v != "" {
			date, _ := time.Parse(shortForm, v)
			month := int(date.Month())
			day, _ := (strconv.ParseInt(strconv.Itoa(date.Day()), 16, 64))
			tail = append(tail, byte(day))
			tail = append(tail, byte(month))
		}

	}

	lng := len(tail)
	rem := lng % 16

	if rem != 0 {
		for i := 0; i <= 16-rem; i++ {
			tail = append(tail, 255)
		}
	}

	pcount := len(tail) / 16

	cnt := 0
	for j := 0; j < pcount; j++ {
		pack := make([]byte, 16)
		copy(pack, tail[cnt:cnt+16])
		cnt += 16
		pack = append(pack, byte(j))
		command := PrepareSetterCommand(netNumber, 16, &pack)
		_, res := PerformCommand(command, portname, timeout, baud, 7)

		if res == false {
			return false, nil
		}
	}

	return true, nil
}
