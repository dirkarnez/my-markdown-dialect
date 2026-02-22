package main

import (
	"fmt"

	"github.com/dirkarnez/my-markdown-dialect/markup"
)

func main() {
	o, _ := markup.Parse("This sentence uses `$` delimiters to show math inline: $\\sqrt{3x-1}+(1+x)^2$<hr>```123```")
	fmt.Println(markup.Finalize(o))
}
