
package loader

import (
  "io/ioutil"
  "encoding/json"
  // "fmt"
  "reflect"
)

// Take in JSON content and parse.
// What you get is a slice with slices and strings inside.
func ParseText(content []byte) []interface{} {
  jsonObject := []interface{}{}
  if err := json.Unmarshal(content, &jsonObject); err != nil {
    panic(err)
  }
  return parse(jsonObject)
}

// Parse JSON from a file.
func ParseFile(filename string) []interface{} {
  content, _ := ioutil.ReadFile(filename)
  return ParseText(content)
}

func parse(obj []interface{}) []interface{} {
  list := &[]interface{}{}
  for _, child := range(obj) {
    value := reflect.ValueOf(child)
    kind := value.Kind()
    switch kind {
    case reflect.Slice:
      length := value.Len()
      item := value.Slice(0, length).Interface()
      inside := parse(item.([]interface{}))
      *list = append(*list, inside)
    case reflect.String:
      *list = append(*list, value.String())
    default:
      panic("got other type from JSON")
    }
  }
  return *list
}