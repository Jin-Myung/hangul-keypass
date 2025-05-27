// main.go
package main

import (
	_ "embed"
	"strings"
	"syscall/js"
)

//go:embed internal/words.txt
var wordText string

var wordList []string

func init() {
	wordList = strings.Split(strings.TrimSpace(wordText), "\n")
}

func generate(this js.Value, args []js.Value) interface{} {
	// TODO: 실제 생성 로직 대체 예정
	password := "tkfrlsk1!"
	origin := "사과나무"

	js.Global().Get("document").Call("getElementById", "result").Set("innerText", password)
	js.Global().Get("document").Call("getElementById", "origin-word").Set("innerText", origin)
	return nil
}

func main() {
	js.Global().Set("generatePassword", js.FuncOf(generate))
	select {}
}
