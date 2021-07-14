package main

import (
	"strings"
	"syscall/js"

	"github.com/notnoobmaster/luautil"
)

func beautify(this js.Value, inputs []js.Value) interface{} {
	inp := inputs[0]
	str := inp.JSValue().String()
	ret, err := luautil.Beautify(strings.NewReader(str))
	if err != nil {
		return str
	}
	return ret
}

func main() {
	c := make(chan bool)
	js.Global().Set("Beautify", js.FuncOf(beautify))
	<-c
}