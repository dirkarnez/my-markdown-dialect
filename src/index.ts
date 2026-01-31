import rehypeSanitize from 'rehype-sanitize'
import rehypeStringify from 'rehype-stringify'
import remarkParse from 'remark-parse'
import remarkRehype from 'remark-rehype'
import {unified} from 'unified'

const file = await unified()
  .use(remarkParse)
  .use(remarkRehype)
  .use(rehypeSanitize)
  .use(rehypeStringify)
  .process('# Hello, Neptune!')

console.log(String(file))


import remarkParse from 'remark-parse'
import remarkStringify from 'remark-stringify'
import {unified} from 'unified'
import {visit} from 'unist-util-visit'

const file = await unified()
  .use(remarkParse)
  .use(myRemarkPluginToIncreaseHeadings)
  .use(remarkStringify)
  .process('# Hi, Saturn!')

console.log(String(file)) // => '## Hi, Saturn!'

function myRemarkPluginToIncreaseHeadings() {
  /**
   * @param {Root} tree
   */
  return function (tree) {
    visit(tree, function (node) {
      if (node.type === 'heading') {
        node.depth++
      }
    })
  }
}
