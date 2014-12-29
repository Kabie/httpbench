package main

import (
  "net/http"
  "fmt"
  "html"
  "encoding/json"
)

type Hello struct {
  Hello string `json:hello`
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    var j Hello
    if r.URL.Path == "/" {
      j = Hello{"world!"}
    } else {
      j = Hello{html.EscapeString(r.URL.Path[1:])}
    }
    s, _ := json.Marshal(j)
    fmt.Fprintf(w, "%s", s)
  })
  http.ListenAndServe(":8080", nil)
}
