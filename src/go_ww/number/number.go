package number

import (
	"fmt"
	"math"
	"unsafe"
)

func LearnNumber() {
	var i8 int8     // 1 个字节
	var i16 int16   //2
	var i32 int32   //4
	var i64 int64   //8
	var ui8 uint8   //1
	var ui16 uint16 //2
	var ui32 uint32 //4
	var ui64 uint64 // 8

	fmt.Printf("%T %dB% %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", ui8, unsafe.Sizeof(ui8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", ui16, unsafe.Sizeof(ui16), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", ui32, unsafe.Sizeof(ui32), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", ui64, unsafe.Sizeof(ui64), math.MinInt8, math.MaxInt8)

	var f64 float64 // 8
	var f32 float32 // 4
	fmt.Printf("%T %dB% %v~%v\n", f64, unsafe.Sizeof(f64), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB% %v~%v\n", f32, unsafe.Sizeof(f32), math.MinInt8, math.MaxInt8)

	var i int // 8
	fmt.Printf("%T %dB% %v~%v\n", i, unsafe.Sizeof(i))
}
