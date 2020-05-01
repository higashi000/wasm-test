package main

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/ikawaha/kagome/tokenizer"
)

func main() {
	document := js.Global().Get("document")
	text := document.Call("getElementById", "text")
	res := document.Call("getElementById", "result")

	t := tokenizer.New()
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
}
