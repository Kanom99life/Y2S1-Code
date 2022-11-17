// ธีรภัทร์ นิลศิริ
// 640510662
// HW01_2
// 204203 Sec 001
package main
import (
	"strings"
	"math"
)

func additiveInverse(x string) (string, int64) {
	//var result int64 = 0
	result := []byte(strings.Repeat("x", len(x)))
	lenX := len(x)-1
	//Traverse binary bits from right to left
	//Find the first 1 bit
	//Reverse every bit after first 1’s
	for i := lenX; i >= 0; i--{
		if x[i] == '1'{
			for j := i; j >= 0; j--{	
			if x[j] == '1'{
				result[j] = '0'	
			}else {
				result[j] = '1'
				}
			}
			result[i] = '1'
			break
		}else{
			result[i] = x[i]
		}
	}
	return string(result[0:]),twosComplToInt(string(result[0:]))	
}

func twosComplToInt(x string) int64 {
	var result int64 = 0
	p := 0
	for i := len(x)-1 ; i >= 0; i--{
		if x[i] == '1'{
			if (i == 0) && (x[i] == '1'){
				result += -(int64(math.Pow(2,float64(p))))
			}else{
				result += int64(math.Pow(2,float64(p)))
				p ++
			}	
		}else{
			p++
		}
	}
	return result
}