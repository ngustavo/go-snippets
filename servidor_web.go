package main

import (
    "fmt"
    "net/http"
    "html"
    "log"
)

func main() {
    http.HandleFunc("/", handler)
    fmt.Println(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<h1>Alo!</h1>")
}

func minimal(){
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
