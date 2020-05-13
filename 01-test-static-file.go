package main

import (
	"net/http"
	"fmt"
)

// 测试静态文件访问
func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server {
		Addr: "0.0.0.0:8001",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s", request.URL.Path[1:])
}