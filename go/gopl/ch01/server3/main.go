package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	// 打印请求头
	for title, value := range r.Header {
		fmt.Fprintf(w, "header %s = %s\n", title, value)
	}

	fmt.Fprintf(w, "Host : %s\n", r.Host)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	// 打印请求参数
	for field, value := range r.Form {
		fmt.Fprintf(w, "From %q = %q\n", field, value)
	}
}
