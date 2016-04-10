package main

import (
  "fmt"
  "regexp"
  "bytes"
)

func main() {
  var logRes = `[2009-11-15T14:13:12][1] -------trying out Elasticsearch
[2009-11-15T14:13:12][2] -------trying out Elasticsearch
[2009-11-15T15:18:12][3] -------trying out Elasticsearch
`
		fmt.Printf("%s\n", replace(logRes, `(\[.*\]){2}`, ""))
}

func replace(source, pattern, target string) string {
	reg := regexp.MustCompile(pattern)
	rep := []byte(target)
	return fmt.Sprintf("%q", reg.ReplaceAll([]byte(source), rep))
}
