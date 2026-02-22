package markup

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func newParser() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			/*html.WithWriter(html.NewWriter()),*/
			html.WithUnsafe(),
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
}

func ParseString(source string) (string, error) {
	var buf bytes.Buffer
	if err := newParser().Convert([]byte(source), &buf); err != nil {
		return "", err
	} else {
		return buf.String(), nil
	}
}

func ParseBytes(source []byte) (string, error) {
	var buf bytes.Buffer
	if err := newParser().Convert(source, &buf); err != nil {
		return "", err
	} else {
		return buf.String(), nil
	}
}
