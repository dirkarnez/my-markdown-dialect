package markup

import (
	"bytes"
	"html/template"
	"log"
)

func Finalize(markdownHTML string) (string, error) {
	// 2. Parse the template content
	// We use template.Must() for simplicity in this example to panic on error
	// otherwise, we would check the error returned by Parse()
	tmpl, err := template.New("hello").Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.0/styles/github-dark.min.css">
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.0/highlight.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.0/languages/matlab.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/highlightjs-line-numbers.js/dist/highlightjs-line-numbers.min.js"></script>
  	<style>
    .hljs-ln-numbers {
        text-align: center;
        color: #ccc;
        border-right: 1px solid #999;
        vertical-align: top;
        padding-right: 5px !important;

        -webkit-touch-callout: none;
        -webkit-user-select: none;
        -khtml-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
    }
    .hljs-ln-code {
        padding-left: 10px !important;
    }

    /* .hljs-ln-code,
    .hljs-ln-numbers {
        line-height: 14px;
    } */

    code {
        white-space: pre-wrap;
        overflow: auto;
    }

	@media print {
		hr {
            page-break-after: always;
            visibility: hidden;
            height: 0;
            padding: 0;
            margin-top: 0;
            margin-bottom: 0;
		}
	}
	</style>
</head>
<body>

<div class="container">
	{{ .Cotent }}
</div>

<script>
document.addEventListener('DOMContentLoaded', (event) => {
  hljs.highlightAll();
  hljs.initLineNumbersOnLoad();
});
</script> 

<script>
(function () {
  const body = document.body.textContent;
  if (body.match(/(?:\$|\\\(|\\\[|\\begin\{.*?})/)) {
    if (!window.MathJax) {
      window.MathJax = {
        tex: {
          inlineMath: {'[+]': [['$', '$']]}
        }
      };
    }
	
    const script = document.createElement('script');
    script.src = 'https://cdn.jsdelivr.net/npm/mathjax@4/tex-chtml.js';
    document.head.appendChild(script);
  }
})();
</script>
</body>
</html>
`)
	if err != nil {
		log.Fatal(err)
	}

	// Data to pass to the template
	data := struct {
		Cotent template.HTML // Use template.HTML for trusted content
	}{
		Cotent: template.HTML(markdownHTML), // This content is marked as safe
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
