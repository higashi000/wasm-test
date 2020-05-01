package main

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/ikawaha/kagome/tokenizer"
)

func main() {
	t := tokenizer.New()
	MorphologicalAnalysis(t)
	go Clear()

	select {}
}

func Clear() {
	clearBtn := js.Global().Get("document").Call("getElementById", "clear")

	cb := js.FuncOf(func(val js.Value, vals []js.Value) interface{} {
		result := js.Global().Get("document").Call("getElementById", "result")
		text := js.Global().Get("document").Call("getElementById", "text")

		result.Set("innerHTML", "")
		text.Set("value", "")

		return nil
	})

	clearBtn.Call("addEventListener", "click", cb)
}

func MorphologicalAnalysis(t tokenizer.Tokenizer) {
	document := js.Global().Get("document")
	text := document.Call("getElementById", "text")
	button := document.Call("getElementById", "button")
	res := document.Call("getElementById", "result")

	cb := js.FuncOf(func(val js.Value, args []js.Value) interface{} {
		tokens := t.Tokenize(text.Get("value").String())

		for _, token := range tokens {
			if token.Class == tokenizer.DUMMY {
				fmt.Printf("%s\n", token.Surface)
				continue
			}
			features := strings.Join(token.Features(), ",")
			fmt.Printf("%s\t%v\n", token.Surface, features)
			features = "<p>" + token.Surface + ": " + features + "</p>"
			res.Call("insertAdjacentHTML", "beforeend", features)
		}
		return nil
	})

	button.Call("addEventListener", "click", cb)
}
