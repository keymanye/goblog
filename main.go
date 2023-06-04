package main

import (
	"fmt"
	"net/http"
)

func Defaulthandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
		fmt.Fprint(w, "客户端请求路径为 "+r.URL.Path)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "你访问的页面没有找到")
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "你现在正在访问about目录")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", Defaulthandler)
	router.HandleFunc("/about", AboutHandler)
	http.ListenAndServe(":3000", router)
}
