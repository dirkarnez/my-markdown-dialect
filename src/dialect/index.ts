// import rehypeSanitize from 'rehype-sanitize'
// import rehypeStringify from 'rehype-stringify'
// import remarkParse from 'remark-parse'
// import remarkRehype from 'remark-rehype'
// import {unified} from 'unified'

// const file = await unified()
//   .use(remarkParse)
//   .use(remarkRehype)
//   .use(rehypeSanitize)
//   .use(rehypeStringify)
//   .process('# Hello, Neptune!')

// console.log(String(file))


import remarkParse from 'remark-parse'
import remarkStringify from 'remark-stringify'
import {unified} from 'unified'
import {visit} from 'unist-util-visit'

export const parseSample = async (str: string): Promise<string> => {
  function myRemarkPluginToIncreaseHeadings() {
    /**
     * @param {Root} tree
     */
    return function (tree: any) {
      visit(tree, function (node) {
        if (node.type === 'heading') {
          node.depth++
        }
      })
    }
  }

  const file = await unified()
    .use(remarkParse)
    .use(myRemarkPluginToIncreaseHeadings)
    .use(remarkStringify)
    .process(str); // '# Hi, Saturn!')
  
  return String(file);
} 


export const parseToHTML = async (str: string): Promise<string> => {
  const file = await unified()
    .use(remarkParse)
    .use(() => (
      (tree: any) => {
        visit(tree, node => {
          if (node.type === 'heading') {
            node.depth++
          }
        })
      }
    ))
    .use(remarkStringify)
    .process(str); // '# Hi, Saturn!')
  
  return String(file);
} 