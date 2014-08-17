
package loader

import (
  "fmt"
  "testing"
  "io/ioutil"
)

func TestLoader(t *testing.T) {
  files, _ := ioutil.ReadDir("json")
  for _, name := range(files) {
    fullname := fmt.Sprintf("json/%s", name.Name())
    result := ParseFile(fullname)
    fmt.Printf("%v\n\n", result)

    content, _ := ioutil.ReadFile(fullname)
    ExampleParseText(content)

  }
}