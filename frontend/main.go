// +build js,wasm

package main

import (
	"syscall/js"

	"github.com/akyoto/wasm-hello/components"
)

var document = js.Global().Get("document")

func main() {
	main := document.Call("getElementById", "main")
	main.Set("innerHTML", components.Hello())
}
