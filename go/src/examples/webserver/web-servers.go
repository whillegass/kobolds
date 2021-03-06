package main

import (
       "fmt"
        "log"
        "net/http"
)

type String string

type Struct struct {
  Greeting string
  Punct    string
  Who      string
}

func (h Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
      fmt.Fprint(w, h)
    }

func (h String) ServeHTTP(
        w http.ResponseWriter,
        r *http.Request) {
        fmt.Fprint(w, h)
}
func main() {
        http.Handle("/string", String("I'm a frayed knot."))
        http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
        err := http.ListenAndServe("localhost:4000", nil)
        if err != nil {
                log.Fatal(err)
        }
}
