// ธีรภัทร์ นิลศิริ
// 640510662
// HW01_1
// 204203 Sec 001

package main
import(
	 "math"
)
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
