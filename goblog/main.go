package main

import (
	"fmt"
	"net/http"
)

func SayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>JiangYe</h1>")
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, fmt.Sprintf("%s", "<h1>Hi  这里是goblog </h1>"))
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, fmt.Sprintf("%s %s", "您正在访问网页是: ", r.URL.Path))
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, fmt.Sprintf("<h1>%s %s %s</h1>", "您正在访问网页: ", r.URL.Path, "不存在"))
	}
}

func main() {
	http.HandleFunc("/", HandlerFunc)
	http.HandleFunc("/sayhello", SayHelloName)
	http.ListenAndServe(":9191", nil)

}
