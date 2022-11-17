// ธีรภัทร์ นิลศิริ
// 640510662
// HW01_3
// 204203 Sec 001
package main

import ("strings"
		"math"
)
func addNSubtract(n1, n2 string, bitLen uint8) (int64, int64) {
	// println("n1 = ", n1)
	// println("n2 = ", n2)
	// println("bitLen = ", bitLen)
	len1 := len(n1)
	len2 := len(n2)
	if len2 > len1{
		len1, len2 = len2, len1
		n1,n2 = n2, n1
	}
	if len1 != len2{
		for i := (len1-len2); i > 0; i--{
			if n2[0] == '1'{
				n2 = string('1') + n2
			}else{
				n2 = string('0') + n2
			}
		}	
	}
	len1 = len(n1)
	len2 = len(n2)
	diff := float64(len1 - int(bitLen))
	diff = math.Abs(diff)
	if int64(diff) != 0{
		for i := int64(diff) ; i > 0; i--{
			if n1[0] == '1'{
				n1 = string('1') + n1
			}else{
				n1 = string('0') +n1
			}
			
			if n2[0] == '1'{
				n2 = string('1') + n2
			}else{
				n2 = string('0') +n2
			}
		}
	}
	// println("diff = ", int64(diff))
	 //println("new n1 = ", n1)
	 //println("new n2 = ", n2)
	return (add(n1, n2, bitLen)), sub(n1,n2,bitLen)
}

func add(n1, n2 string, bitLen uint8) int64{
	result := []byte(strings.Repeat("0", int(bitLen)))
	
	lenX := len(n1)-1
	k := int(bitLen)-1
	carry := 0
	for i := lenX; i >= 0; i--{
		currDigit := 0 + carry
		currDigit += int(n1[i]) - int('0')
		currDigit += int(n2[i]) - int('0')
		if currDigit == 2 {
			currDigit = 0
			carry = 1
		}else if currDigit == 3{
			currDigit = 1
			carry =1
		}else{
			carry = 0
		}
		result[k] = byte(currDigit%10 + int('0')) 
		if(k == 0){
			result[k] = byte(currDigit + int('0'))
			k--
			break
		}
		k--
		if(carry == 1) && (i == 0){
			result[k] = byte(carry + int('0'))
			k--
		}
	}
	//println("result = ",string(result))
	return twosComplToInt(string(result[:]))
}

func sub(n1, n2 string, bitLen uint8) int64{
	n2 = additiveInverse(n2)
	//println("n2 = ",n2)
	return add(n1,n2,bitLen)
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

func additiveInverse(x string) string {
	result := []byte(strings.Repeat("x", len(x)))
	lenX := len(x)-1
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
	return string(result[0:])
}



