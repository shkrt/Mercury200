package util

import (
	"fmt"
	"mercury200/crc16"
	"strconv"
)

func SplitEvery(source string, step int) []string {
	result := make([]string, len(source)/step)
	var res string
	pos := 0
	for i, v := range source {
		res += string(v)
		if (i+1)%step == 0 {
			result[pos] = res
			pos++
			res = ""
		}
	}

	return result
}

//"266608" => [4 17 112]
func NetNumToArr(netNumber string) []byte {
	res := make([]byte, 3)
	i, _ := strconv.ParseInt(netNumber, 0, 64)
	r := fmt.Sprintf("%06x", i)
	x := SplitEvery(r, 2)

	for ind, v := range x {
		var s, _ = strconv.ParseInt(v, 16, 64)
		res[ind] = byte(s)
	}

	return res
}

//[00 04 17 112 28] => [50 EB]
func GetCrcBytes(command []byte) []byte {
	res := make([]byte, 2)
	crc16 := crc16.CheckSum(command)
	r := fmt.Sprintf("%04x", crc16)
	x := SplitEvery(r, 2)

	for ind, v := range x {
		var s, _ = strconv.ParseInt(v, 16, 64)
		res[ind] = byte(s)
	}

	tmp := res[0]
	res[0] = res[1]
	res[1] = tmp

	return res
}

func CheckCrc(response []byte, respLen int) bool {
	resp := response[0 : respLen-2]
	respCrc := response[respLen-2 : respLen]
	crcBytes := GetCrcBytes(resp)
	if SliceEq(crcBytes, respCrc) {
		return true
	}
	return false
}

func SliceEq(a, b []byte) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
