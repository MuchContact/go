package main

import (
  "fmt"
  "html/template"
  "bytes"
)

func main() {

  const templateText = `
  Input: {{ .cpus_elasticsearch }}
  `
  t, _ := template.New("monitor").Parse(templateText) // Parse template file.
  var doc bytes.Buffer
  conf := make(map[string]string, 0)
  conf["cpus_elasticsearch"]  = "0.2"
  err := t.Execute(&doc, conf)
  if(err != nil){
    fmt.Print(err)
  }
  resp := doc.String()
  fmt.Print(resp)
}
