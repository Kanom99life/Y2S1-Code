// ธีรภัทร์ นิลศิริ
// 640510662
// HW02_1
// 204203 Sec 001

package main

import (
	"strconv"
	"strings"
)

func ipv4Encode(ipString string) uint32 {
	var result uint32 = 0
	var x1,x2,x3,x4 uint32
	numSet := strings.Split(ipString, ".")

	firstB,_ := strconv.Atoi(string(numSet[0]))
	secondB,_ := strconv.Atoi(string(numSet[1]))
	thirdB,_ := strconv.Atoi(string(numSet[2]))
	fourthB,_ := strconv.Atoi(string(numSet[3]))
	
	x1 = uint32(firstB) << 24
	x2 = uint32(secondB) << 16
	x3 = uint32(thirdB) << 8
	x4 = uint32(fourthB)
	result = ( x1 | x2 | x3 | x4 )
	return result
}

func ipv4Decode(ipUint uint32) string {
	var result = ""
	shiftRange := 24
	var rShift, lShift uint32
	for i := 0; i < 4; i++{
		rShift = ipUint >> shiftRange
		lShift = rShift << shiftRange
		sub := ipUint - lShift
		ipUint = sub
		shiftRange -= 8
		
		if(i != 3){
			result += strconv.Itoa(int(rShift)) + "."
		}else{
			result += strconv.Itoa(int(rShift)) 
		}
	}
	return result
}
