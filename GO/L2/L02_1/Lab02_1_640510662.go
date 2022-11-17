// ธีรภัทร์ นิลศิริ
// 640510662
// Lab02_1
// 204203 Sec 001

package main

func powerOfTwo(n uint64) bool {
	return n & (n-1) == 0
}