// ธีรภัทร์ นิลศิริ
// 640510662
// Lab02_2
// 204203 Sec 001

package main

import (
	"strings"
	"math"
)

const MAX = 73 

func baseNAddition(r1, r2 string, base int) string {
	// Split a strings from "."
	splitN1 := strings.Split(r1, ".")
	splitN2 := strings.Split(r2, ".")

	//split[0] == front value
	//split[1] == back value

	//Seperate numbers after decimal ponit.
	var backR1, backR2 string
	var lenBR1, lenBR2 int
	//Check if there is a decimal.
	if len(splitN1) == 1 { 	//If there is no numbers after the decimal point, it will equal to one.
		backR1 = "0"		//Set its value and len to zero.
		lenBR1 = 0
	}else{
		backR1 = splitN1[1]	
		lenBR1 = len(backR1)
	}
	if len(splitN2) == 1 {
		backR2 = "0"
		lenBR2 = 0
	}else{
		
		backR2 = splitN2[1]
		lenBR2 = len(backR2)
	}
	//End of seperate decimal point.
	 
	maxLen := math.Max(float64(lenBR2), float64(lenBR1)) //Find the maximum len.
	
	//Make the len of both numbers equal
	newBr1 := backR1
	newBr2 := backR2
	if lenBR2 > lenBR1{
		newBr1 = eqBackLen(backR1, backR2)
	}else{
		newBr2 = eqBackLen(backR1, backR2)
	}

	//Assign new value of each numbers after making the len of both numbers equal.
	r1 = newBr1 
	r2 = newBr2
	//Calculate the value of decimals
	backValue := addition(r1, r2, base)
	backLen := len(backValue) 

	//Check carry number after sum of two decimals.
	var overLen string
	if backLen > int(maxLen) && backValue != "0"{ //If there is carry number the first digit will use to sum with an integer.
		overLen = string(backValue[0])
		backValue = backValue[1:]				//Keep only a decimals.
	}
	//****************End of calculate decimal.****************//

	//Seperate numbers before a decimal ponit.
	frontR1 := splitN1[0]
	frontR2 := splitN2[0]
	frontValue := frontCal(frontR1, frontR2, base) //Call the integer caculater function.
	
	//If there is a carry number it will be use to sum with an integer.
	if backLen > int(maxLen){
		frontValue = addition(frontValue,overLen,base)
	}
	//****************End of calculate integer.****************//

	//Output
	var finalValue string
	if backValue != "0"{
		finalValue = frontValue + "." + backValue //If there is an decimal.
	}else{
		finalValue = frontValue //If there is no decimal.
	}
	return finalValue
	//****************MAIN****************//
}

//Function for sum numbers
func addition(n1, n2 string, base int) string { 
	result := []byte(strings.Repeat("x", MAX))
	len1 := len(n1)
	len2 := len(n2)
	carry := 0
	i, j, k := len1-1, len2-1, MAX-1
	for ; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		temp := carry
		carry = 0
			if i >= 0 {
				temp += int(n1[i]) - int('0')
			}
			if j >= 0 {
				temp += int(n2[j]) - int('0')
			}
			result[k] = byte((temp % base) + int('0'))
			carry = temp / base
	}
	if carry > 0 {
		result[k] = '1'
	} else {
		k++
	}
	return string(result[k:])
}

//Function to make both len of numbers equal.
func eqBackLen(n1, n2 string) string{
	len1 := len(n1)
	len2 := len(n2)
	if len2 > len1{
		len1, len2 = len2, len1
		n1,n2 = n2, n1
	}
	if len1 != len2{
	for i := (len1-len2); i > 0; i--{
		n2 = n2 + "0"
		}
	}
	return n2
}

//Functions for sum integers.
func frontCal(f1, f2 string, base int)string{
	frontValue := addition(f1,f2,base)
	return frontValue
}