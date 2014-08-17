
package loader

func ExampleParseText(content []byte) {
  result := ParseText(content)
  var handleAst func([]interface{})

  handleAst = func(ast []interface{}) {
    for _, child := range(ast) {
      if expr, ok := child.([]interface{}); ok {
        handleAst(expr)
      } else if token, ok := child.(string); ok {
        print(token)
        print(" ")
      } else {
        panic("other type appeared")
      }
    }
  }
  handleAst(result)
}