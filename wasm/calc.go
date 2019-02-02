package main // import "calc"

import (
	"strconv"
	"syscall/js"
)

func add(i []js.Value) {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
}

func subtract(i []js.Value) {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
	js.Global().Set("subtract", js.NewCallback(subtract))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")

	registerCallbacks()
	<-c
}
