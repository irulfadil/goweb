package main

import (
	"fmt"
	"syscall/js"

	_ "github.com/go-sql-driver/mysql"
)

// login exposed function to JS interface{}
func login(this js.Value, args []js.Value) interface{} {
	// Get the JS objects here
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return js.Global().Call("eval", `Swal.fire("Oops!, Error", "Unable to get document object", "error");`)
	}
	username := jsDoc.Call("getElementById", "username")
	if !username.Truthy() {
		return js.Global().Call("eval", `Swal.fire("Oops!, Error", "Unable to get username", "error");`)
	}
	return nil
}

func exposeGoFunc() {
	// Start exposing the following Go functions to JS client side
	js.Global().Set("login", js.FuncOf(login))
}

func main() {
	fmt.Println("Welcome to WASM cc: irulfadil")
	c := make(chan bool, 1)

	// Initializes all your exposable Go's functions to JS
	exposeGoFunc()
	<-c
}
