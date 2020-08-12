package bluetooth

import "syscall/js"

var (
	navigator = js.Global().Get("navigator")
	bluetooth = navigator.Get("bluetooth")
	console   = js.Global().Get("console")
)
