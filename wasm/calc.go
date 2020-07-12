package main // import "calc.wasm"

import (
	"strconv"
	"syscall/js"
)

var jb = js.Global()

func add(this js.Value, i []js.Value) interface{} {
	value1 := i[0].String()
	value2 := i[1].String()
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	if value1 == "" {
		int1 = 0
	}
	if value2 == "" {
		int2 = 0
	}

	result := int1 + int2

	return result
}

func sub(this js.Value, i []js.Value) interface{} {
	value1 := i[0].String()
	value2 := i[1].String()
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	if value1 == "" {
		int1 = 0
	}
	if value2 == "" {
		int2 = 0
	}

	result := int1 - int2

	return result
}

func multi(this js.Value, i []js.Value) interface{} {
	value1 := i[0].String()
	value2 := i[1].String()
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	if value1 == "" {
		int1 = 0
	}
	if value2 == "" {
		int2 = 0
	}

	result := int1 * int2

	return result
}

func divi(this js.Value, i []js.Value) interface{} {
	value1 := i[0].String()
	value2 := i[1].String()
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	if value2 == "" || value2 == "0" {
		return "infinite"
	}

	result := int1 / int2

	return result
}

func main() {
	c := make(chan struct{}, 0)

	jb.Set("waAdd", js.FuncOf(add))
	jb.Set("waSub", js.FuncOf(sub))
	jb.Set("waMulti", js.FuncOf(multi))
	jb.Set("waDivi", js.FuncOf(divi))

	println("Go Web Assembly Ready")

	<-c
}
