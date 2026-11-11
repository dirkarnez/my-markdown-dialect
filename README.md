my-markdown-dialect
===================
<kbd>[**dirkarnez/my-markdown-dialect-essay-template**](https://github.com/dirkarnez/my-markdown-dialect-essay-template)</kbd><br>


### Elements to support
- [ ] `import` code files
- [ ] `<div>`-based flex-box columns (with start / end n columns style)
- [ ] code diff left-right view
- [ ] css grid table
- [x] code blocks with line number
- [x] math blocks
    - 2 `$` for centered `$$\sum_{i=1}^n i = \frac{n(n+1)}{2}$$`
    - 1 `$` for left `$$\sum_{i=1}^n i = \frac{n(n+1)}{2}$$`
- [ ] `import` [JSXGraph](https://jsxgraph.uni-bayreuth.de/home/)
- [x] explicit newlines
  - ```html
    1<br>2
    ```
- [x] explicit newlines
  - ```html
    1<hr>2
    ```
- [x] image and caption
  - ```html
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/Wiki-Loves-Love-logo.svg/120px-Wiki-Loves-Love-logo.svg.png" style="width: 100%">
    ```
- [ ] quotes
- [ ] footnotes
- [x] Force center
  - ```html
    <div style="text-align: center;">
        123
    </div>
    ```
- [x] Left-aligned calculation steps
  - [Aligning Equations in mathjax - Write down the problem. Think real hard. Write down the solution.](https://marquis08.github.io/mathjax/minimalmistakes/align-equations-in-mathjax/)
      - ```markdown
        $$
        \begin{aligned}
        \sum_{i=0}^{\infty} i^2 & = \sqrt{\frac{73^2-1}{12^2}} \\
         & = \sqrt{\frac{73^2}{12^2}\cdot\frac{73^2-1}{73^2}} \\ 
         & = \sqrt{\frac{73^2}{12^2}}\sqrt{\frac{73^2-1}{73^2}} \\
         & = \frac{73}{12}\sqrt{1 - \frac{1}{73^2}} \\ 
         & \approx \frac{73}{12}\left(1 - \frac{1}{2\cdot73^2}\right)
        \end{aligned}
        $$
        ```
- [ ] Letter
    - ```html
      <!DOCTYPE html>
      <html lang="zh-Hant">
      <head>
      <style>
        .p1 {
          font-family: "Times New Roman", Times, serif;
        }

      </style>
      </head>
      <body>
       <p class="p1">This is a paragraph, shown in the Times New Roman font 一仍蛛焜.</p>
      </body>
      </html>
      ```





### TODOs
- [ ] Headless chrome pdf (my desktop chrome is messed up)
  - https://github.com/chromedp/cdproto/blob/master/page/page.go
  - https://github.com/chromedp/examples/blob/master/pdf/main.go
  - [dirkarnez/savepdf](https://github.com/dirkarnez/savepdf)
  - https://github.com/chromedp/examples/tree/master/remote
  - https://github.com/rahul-bhave/headless-chrome-github-action
- [ ] Linux scripts
- [ ] WebAssembly
- [ ] default file name and specified file name

### Tutorials
- [moyen-blog/goldmark-extensions: Extensions for the GoldMark Markdown parser used by Moyen](https://github.com/moyen-blog/goldmark-extensions)

### Addons
- [remarkjs/remark-toc: plugin to generate a table of contents (TOC)](https://github.com/remarkjs/remark-toc)

### Preview and printing
- [dirkarnez/html-repl](https://github.com/dirkarnez/html-repl)
    - ```js
      function printIframeContent() {
        var iframe = document.getElementById("print-iframe");
        
        // Wait for the content to load (especially crucial if it's a PDF file)
        iframe.onload = function() {
            var contentWindow = iframe.contentWindow;
            contentWindow.focus(); // Set focus to the iframe
            contentWindow.print(); // Open the print dialog
        };
        
        // If the content is already loaded, you can call print() directly
        if (iframe.contentWindow.document.readyState === 'complete') {
            iframe.onload(); // Trigger the onload logic manually
        }
        }
     ```

### Go
- [Advanced markdown processing in Go](https://blog.kowalczyk.info/article/cxn3/advanced-markdown-processing-in-go.html)
- [gomarkdown/markdown: markdown parser and HTML renderer for Go](https://github.com/gomarkdown/markdown)

### Reference
- [liffiton/code2tex: Convert source code to a LaTeX document / PDF with syntax highlighting and line numbers.](https://github.com/liffiton/code2tex)
