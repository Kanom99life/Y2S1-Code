// ธีรภัทร์ นิลศิริ
// 640510662
// HW02_2
// 204203 Sec 001

package main
import(
	"strings"
	"strconv"
	"math"
	//"fmt"
)

func roundToEven(x string, bPlace uint8) string {
	splitX := strings.Split(x, ".")
	
	//sepInt:= splitX[0] //ตัวแปรเก็บค่าหน้าจุดทศนิยม

	sepInt := splitX[0]
	lenSepInt := len(sepInt)
	var sepFloat, result string //สร้างตัวแปรเก็บค่าหลังจุดทศนิยมถ้ามีจะเก็บจาก splitX[1]
	var lensepFloat int
	if len(splitX) > 1{
		sepFloat = splitX [1]
		lensepFloat = len(sepFloat)
	}
	//println(lensepFloat)
	
	if(bPlace == 0) && (lensepFloat == 0 ){ // กรณที่เป็นจำนวนเต็มและbPlace เป็น 0 คืนค่า input ไปเลย
		result = x
		return result
	}
	if(int(bPlace) == lensepFloat){ // กรณีที่จำนวนทศนิยมและbPlaceมีค่าเท่ากันคืนค่า input ไปเลย
		result = x
		return result
	}

	if(lensepFloat < int(bPlace)){//กรณีที่inputเป็นจำนวนเต็มหรือมีทศนิยมแต่bPlaceมากกว่า1 เติมทศนิยมเพิ่มให้เท่าbPlace 
		for i := (int(bPlace)-lensepFloat); i > 0; i--{ 
			sepFloat = sepFloat + "0"
		}
		result = sepInt + "." + sepFloat
		return result
	}
	
	if(lensepFloat > int(bPlace)){//กรณีที่ตัดให้เหลือตำแหน่งทศนิยมเท่ากับค่าbPlace
		if(lensepFloat == 1){ //กรณีหลังทศนิยมมี 1 ตัวแล้วจะตัดให้เหลือทศนิยม 0 ตำแหน่ง
			cutFloat := sepFloat[:int(bPlace)+1] //ทศนิยมตัวแรก

			//ถ้าตัวหน้าจุดทศนิยมเป็น 0 และ เลขหลังจุดทศนิยมที่มีตัวเดียวมีค่าเป็น 1 หรือ 0 จะไม่มีการปัดเลขขึ้น
			if (sepInt[lenSepInt-1] == '0') && (cutFloat == "1" || cutFloat == "0"){
				result = sepInt
				return result
			}

			if (sepInt[lenSepInt-1] == '1') && (cutFloat == "1") { //ถ้าตัวหน้าจุดทศนิยมเป็น 1 และ เลขหลังจุดทศนิยมที่มีตัวเดียวมีค่าเป็น 1 จะปัดเลขขึ้น
				roundCut := addition(sepInt, string(cutFloat), lenSepInt)
				result = roundCut
				return result
			}else if (sepInt[lenSepInt-1] == '1') && (cutFloat == "0"){ //ถ้าตัวหน้าจุดทศนิยมเป็น 1 และ เลขหลังจุดทศนิยมที่มีตัวเดียวมีค่าเป็น 0 จะไม่ปัดเลข
				result = sepInt
				return result
			}		

		}else if (lensepFloat > 1){//กรณีหลังทศนิยมมี n ตัวแล้วจะตัดให้เหลือทศนิยม ที่น้อยกว่า n ตำแหน่ง
			//กรณีที่bPlace == 0
			if(int(bPlace) == 0){
				afterCutFloat0 := sepFloat[int(bPlace):] //เก็บค่าจำนวนทศนิยมที่ตัดออกไปหลังตำแหน่ง bPlace
				if (sepInt[len(sepInt)-1] == '0'){ // กรณีที่เลขหลังสุดของตัวหน้าทศนิยมเป็น 0
					if (sumAfterFloatCut(afterCutFloat0)){ //ถ้าเลขทศนิยมมีค่า > 0.5 ปัดเลขขึ้น
						sepInt = addition(sepInt, "1", len(sepInt))
						result = sepInt
						return sepInt
					}else{ //ถ้าเลขทศนิยมมีค่า <= 0.5 ไม่มีการปัดเลข
						return sepInt
					}
				}else{ // กรณีที่เลขหลังสุดของตัวหน้าทศนิยมเป็น 1
					
					if (sepFloat[0] == '1'){ //ถ้าเลขหลังทศนิยมเป็น 1 หรือ มีค่า >= 0.5 จะปัดขึ้น
						sepInt = addition(sepInt, "1", len(sepInt))
						return sepInt
					}else{
						return sepInt
					}

				}
			}

			cutFloat := sepFloat[:int(bPlace)] // เก็บค่าจำนวนทศนิยมให้เท่ากับจำนวน bPlace
			afterCutFloat := sepFloat[int(bPlace):] // เก็บค่าจำนวนทศนิยมที่ตัดออกไปหลังตำแหน่ง bPlace
			//fmt.Println("cutFloat = ",cutFloat)
			//fmt.Println("afterCutFloat = ",afterCutFloat)
			//print(sumAfterFloatCut(string(afterCutFloat)))

			
			lastCutFloat := cutFloat[int(bPlace)-1] //เก็บค่าตำแหน่งสุดท้ายของทศนิยมที่ตัด
			//fmt.Println("lastCutFloat = ",string(lastCutFloat))
			//เช็คว่าต้องปัดเลขขึ้นไหมถ้าตำแหน่งทศนิยมตัวสุดท้ายเป็นเลข 0
			if (lastCutFloat == '0'){
				if(sumAfterFloatCut(afterCutFloat)){ //ถ้าเลขหลังจุดทศนิยมมีค่ามากกว่า 0.5 จะปัดเลขขึ้น
					roundCutFloat := addition(string(cutFloat), "1", len(cutFloat))
					//fmt.Println("RCD0 = ", roundCutFloat)
					result = sepInt + "." + roundCutFloat
					return result
				}else{
					result = sepInt + "." + cutFloat
					return result
				}

			}else if(lastCutFloat == '1'){//ตำแหน่งทศนิยมตัวสุดท้ายเป็นเลข 1 
				//println("lastCutFloat == 1")
				//fmt.Println("afterCutFloat[0] == ", string(afterCutFloat[0]))
				if(afterCutFloat[0] == '1' ){ //ถ้าตัวเลขตำแหน่งแรกของตัวที่ตัดไปเป็น 1 จะปัดทศนิยมขึ่น
					roundCutFloat := addition(string(cutFloat), "1", len(cutFloat))
					//fmt.Println("RCD1 = ", roundCutFloat)

					//กรณีที่ทดเลขแล้วมีค่าเกินให้เอาไปบวกกับเลขหน้าจุดทศนิยม
					if( len(roundCutFloat) > len(cutFloat) ){
						overLen := string(roundCutFloat[0])
						add2Int := addition(sepInt, string(overLen), len(sepInt))
						roundCutFloat = roundCutFloat[1:]

						result = add2Int + "." + roundCutFloat
						return result

					}else{//กรณีที่ทดเลขแล้วมีค่าไม่เกิน
						result = sepInt + "." + roundCutFloat
						return result
					}
				}else{ //ถ้าตัวเลขตำแหน่งแรกของตัวที่ตัดไปเป็น 0 จะไม่มีการปัดทศนิยม
					result = sepInt + "." + cutFloat
					return result
				}
			}					
		}
	}

	return result
}

func sumAfterFloatCut(afCut string) bool { //function ใช้ในการเช็คว่าตัวเลขหลังทศนิยมที่ตัดมีค่ามากว่า 0.5 หรือไม่
	lenAfCut := len(afCut) 
	var curr float64 // เก็บค่าผลบวก
	power := -1 //เลขยกกำลังหลังทศนิยมจะเริ่มที่ -1
	for i := 0; i < lenAfCut; i++{
		num, _ := strconv.Atoi(string(afCut[i]))
		curr += float64(num) * math.Pow(2,float64(power))
		power --	
	}

	//fmt.Println("curr = ",curr)
	if(curr > 0.5){ //ถ้าเลขหลังทศนิยมมีค่า > 0.5 ปัดเลขขึ้น
		return true
	}else{ //ถ้าเลขหลังทศนิยมมีค่า <= 0.5 ไม่ปัดเลข
		return false
	}
}


func addition(n1, n2 string, bPlace int) string { 
	result := []byte(strings.Repeat("x", bPlace+1))
	len1 := len(n1)
	len2 := len(n2)
	if len1 > len2{
		for l := len1-1 ; l > 0; l--{
			n2 = "0" + n2
		}
	}
	len2 = len(n2)
	
	carry := 0
	i, j, k := len1-1, len2-1, bPlace
	for ; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		temp := carry
		carry = 0
			if i >= 0 {
				temp += int(n1[i]) - int('0')
			}
			
			if j >= 0 {
				temp += int(n2[j]) - int('0')
			}
			
			result[k] = byte((temp % 2) + int('0'))
		
			carry = temp / 2		
	}
	
	if carry > 0 {
		result[k] = '1'
	} else {
		k++
	}
	return string(result[k:])
}