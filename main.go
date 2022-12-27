package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	// dataType()
	// floatType()
	// complexType()
	// constant()
	pointer()
}

// dataType
func dataType() {
	// Cách khai báo biến (Nếu không có giá trị khởi tạo thì mặc định = 0)

	//Khai báo 1 biến
	// var intType = 3
	// fmt.Println(intType)

	//Khai báo nhiều biến
	// var width, height = 100, 50
	// fmt.Println("Width = ", width, "Height =", height)

	// var (
	// 	stringType = "String"
	// 	intType2   = 3
	// )
	// fmt.Println(stringType, intType2)

	//Khai báo nhanh (Phải gán giá trị) và phải có ít nhất 1 biến mới (chưa được khai báo trước đó)
	// name, age := "Huy", 21
	// fmt.Println("my name is", name, "age is", age)

}

func intType() {
	//uint (Chỉ giá trị dương), int(có giá trị âm)
	//uint8 : máy tính cần 8 bit đê biểu diễn số nguyên
	// 	KIỂU	GIỚI HẠN
	// uint8	0 – 255
	// uint16	0 – 65535
	// uint32	0 – 4294967295
	// uint64	0 – 18446744073709551615
	// int8	-128 – 127
	// int16	-32768 – 32767
	// int32	-2147483648 – 2147483647
	// int64	-9223372036854775808 – 9223372036854775807
}

func floatType() {
	//Cũng có 2 kiểu float32 và float64 có giới hạn riêng của nó

	//EX
	// a, b := 5.67, 8.97
	// fmt.Printf("Kiểu dữ liệu của a là %T và của b là %T\n", a, b)
	// sum := a + b
	// sub := a - b
	// fmt.Println("Tổng a và b:", sum)
	// fmt.Println("Hiệu a và b:", sub)
}

func complexType() {
	//Có 2 kiểu số phức là complex64 và complex128
	// Complex64: Số phức có phần thực float32 và phần ảo
	// Complex128: Số phức có phần thực float64 và phần ảo

	//Có 2 cách tạo số phức
	// 1 Số phức được tạo bằng một hàm Complex để xây dựng phần thực và phần ảo:
	//	func complex(r, i FloatType) ComplexType
	// 2: Được tạo bằng cú pháp viết tắt:
	//	c := 6 + 7i

	//Khi khai báo nên khai phần thực và phần ảo cùng kiểu dữ liệu

	//ex
	// c1 := complex(3, 3)
	// c2 := 3 + 3i
	// fmt.Println(c1 + c2)
}

func stringType() {
	// first := "Naveen"
	// last := "Ramanathan"
	// name := first + " " + last
	// fmt.Println("My name is", name)
}

func squeezeStyle() {
	//Ngôn ngữ Go rất nghiêm ngặt và chặt chẽ, nên chúng không cho phép tự động chuyển đổi (ép kiểu) kiểu dữ liệu.
	// Ví dụ muốn cộng 2 biến lại thì 2 biến đó phải cùng kiểu dữ liệu
	// Cú pháp: T(v) là cú pháp để chuyển đổi một kiểu dữ liệu của v thành T
	//Ex
	// i := 55      //int
	// j := 67.8    //float64
	// sum := i + int(j) //j is được ép kiểu thành int
	// fmt.Println(sum)
}

func constant() {
	//Từ khoá const là khai báo hằng số và hằng số thì không thay đổi được

	//Tuy nhiên với cách khai báo như này thì hằng số const tự động hiểu
	//Giá trị của 1 hằng số cần được biết ngay tại thời gian biên dịch
	// 	var a int 50
	// var b string = "I love Go"

	//String constant (Hằng số chuỗi)
	var name = "Sam"
	fmt.Printf("type %T value %v", name, name) // type string value Sam
	//Việc gán 1 giá trị có kiểu dữ liệu khác cho 1 biến không được chấp nhận trong Go.

	//Hằng số boolean
	// const trueConst = true
	// type myBool bool
	// var defaultBool = trueConst //chạy bình thường
	// var customBool myBool = trueConst //chạy bình thường
	// defaultBool = customBool //lỗi do 2 biến có kiểu khác nhau

	//Hằng số Số (numeric constant)
	// const a = 5
	// var intVar int = a
	// var int32Var int32 = a
	// var float64Var float64 = a
	// var complex64Var complex64 = a
	// fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)
	//a có thể gán cho bất cứ thằng nào vì a được hiểu là tất cả trong này

	//Biểu thức số học
	//Numeric constant có thể được sử dụng thoaỉ mái trong các biểu thức, và kiểu của nó chỉ cần khi nó được gán cho 1 biến khác, hoặc sử dụng ở những đoạn code yêu cầu trả về kiểu dữ liệu của nó.
}

func pointer() {
	// var p *int
	// fmt.Println(p)

	b := 25
	a := &b

	fmt.Println(a)
	fmt.Println(b)
}
