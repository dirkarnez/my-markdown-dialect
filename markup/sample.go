package markup

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// 1. AST NODE
type HighlightNode struct {
	ast.BaseInline
}

func (n *HighlightNode) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

var KindHighlight = ast.NewNodeKind("Highlight")

func (n *HighlightNode) Kind() ast.NodeKind {
	return KindHighlight
}

// 2. PARSER
type highlightParser struct{}

func (s *highlightParser) Trigger() []byte {
	return []byte{'='}
}

func (s *highlightParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, _ := block.PeekLine()
	if len(line) < 3 || line[0] != '=' || line[1] != '=' {
		return nil
	}
	// Look for closing ==
	idx := bytes.Index(line[2:], []byte("=="))
	if idx <= 0 {
		return nil
	}

	block.Advance(2) // skip opening ==
	node := &HighlightNode{}
	// Process content between == and ==
	content := line[2 : 2+idx]
	node.AppendChild(node, ast.NewTextSegment(text.NewSegment(block.PreciseCursor(), block.PreciseCursor()+len(content))))
	block.Advance(idx + 2) // skip content and closing ==

	return node
}

// 3. RENDERER
type highlightHTMLRenderer struct{}

func (r *highlightHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindHighlight, r.renderHighlight)
}

func (r *highlightHTMLRenderer) renderHighlight(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		w.WriteString("<mark>")
	} else {
		w.WriteString("</mark>")
	}
	return ast.WalkContinue, nil
}

// 4. THE EXTENSION
type highlightExtension struct{}

func (e *highlightExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(&highlightParser{}, 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&highlightHTMLRenderer{}, 500),
	))
}

func main() {
	md := goldmark.New(goldmark.WithExtensions(&highlightExtension{}))
	var buf bytes.Buffer
	source := []byte("This is ==important== text.")

	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}

	fmt.Print(buf.String())
}
