package main

import (
	"os"

	"github.com/dirkarnez/my-markdown-dialect/markup"
)

var (
	title string
	fileName string
)

func main() {
	flag.StringVar(&title, "title", "", "PDF title")
	flag.StringVar(&fileName, "fileName", "", "file name of the input Markdown file")
	flag.Parse()

	var content []byte
	var err error
	
	if len(fileName) < 1 {
		content, err = os.ReadFile("Report.md")
	} else {
		content, err = os.ReadFile(fileName)
	}
	
	if err != nil {
		panic(err)
	}
	
	// o, _ := markup.Parse("This sentence uses `$` delimiters to show math inline: $\\sqrt{3x-1}+(1+x)^2$<hr>```123```")
	o, _ := markup.ParseBytes(title, content)
	finalHTML, err := markup.Finalize(o)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("Report.html", []byte(finalHTML), 0644)
	if err != nil {
		panic(err)
	}
}
