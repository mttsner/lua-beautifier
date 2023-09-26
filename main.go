package main

import (
	"strings"
	"syscall/js"

	"github.com/notnoobmaster/luautil/parse"
)

func beautify(this js.Value, inputs []js.Value) interface{} {
	str := inputs[0].String()
	chunk, err := parse.Parse(strings.NewReader(str), "")
	if err != nil {
		return str
	}
	return chunk.String()
}

func main() {
	c := make(chan bool)
	js.Global().Set("Beautify", js.FuncOf(beautify))
	<-c
}
