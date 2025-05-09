package main

import "fmt"

func main() {
	var iniInt32 int32 = 64
	var iniInt16 int16 = int16(iniInt32)
	var iniInt64 int64 = int64(iniInt16)

	// fmt.Println(iniInt64)
	var iniString string = string(iniInt64)
	fmt.Println(iniString)
}