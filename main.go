package main

import (
	"os"

	"github.com/dirkarnez/my-markdown-dialect/markup"
)

func main() {
	content, err := os.ReadFile("Report.md")
	if err != nil {
		panic(err)
	}
	// o, _ := markup.Parse("This sentence uses `$` delimiters to show math inline: $\\sqrt{3x-1}+(1+x)^2$<hr>```123```")
	o, _ := markup.ParseBytes(content)
	finalHTML, err := markup.Finalize(o)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output.html", []byte(finalHTML), 0644)
	if err != nil {
		panic(err)
	}
}
