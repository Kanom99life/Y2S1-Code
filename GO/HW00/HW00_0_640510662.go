// ธีรภัทร์ นิลศิริ
// 640510662
// HW00_0
// Problem A
// 204203 Sec 001

package main

func factorial(num1 int8) int64 {
	var result int64 = 1
	if num1 <= 1{
		return 1
	}
	result = int64(num1) * factorial(num1 - 1)
	return result
}