package main

import (
	"flag"
	"os"
	"strings"
	
	"github.com/dirkarnez/my-markdown-dialect/markup"
)

var (
	title string
	fileName string
	outputFileName string
)

func main() {
	flag.StringVar(&title, "title", "", "PDF title")
	flag.StringVar(&fileName, "fileName", "", "file name of the input Markdown file")
	flag.StringVar(&outputFileName, "outputFileName", "", "output file name of the HTML file")
	flag.Parse()

	var content []byte
	var err error
	
	if len(fileName) < 1 {
		content, err = os.ReadFile("Report.md")
	} else {
		content, err = os.ReadFile(fileName)
	}

	if len(outputFileName) < 1 {
		outputFileName = "Report.html"
	} else {
		if !strings.HasSuffix(outputFileName, ".html") {
			outputFileName = outputFileName + ".html"
		}
	}
	
	if err != nil {
		panic(err)
	}
	
	// o, _ := markup.Parse("This sentence uses `$` delimiters to show math inline: $\\sqrt{3x-1}+(1+x)^2$<hr>```123```")
	o, _ := markup.ParseBytes(content)
	finalHTML, err := markup.Finalize(title, o)
	if err != nil {
		panic(err)
	}
	
	err = os.WriteFile(outputFileName, []byte(finalHTML), 0644)
	if err != nil {
		panic(err)
	}
}
