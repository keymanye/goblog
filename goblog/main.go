package main

import (
	"fmt"
	"net/http"
	"strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, fmt.Sprintf("%s", "<h1>Hello, 欢迎来到 goblog！</h1>"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "您正在访问网页走丢了 :( "+"<p> 有问题联系我们</p>")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "这是一个blog示例，有问题请联系："+"<a href='mailto:518@msn.cn'>518@msn.cn</a>")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)
	//文章详情
	router.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.SplitN(r.URL.Path, "/", 3)[2]
		fmt.Fprint(w, "文章 ID："+id)
	})

	// 列表 or 创建
	router.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w, "访问文章列表")
		case "POST":
			fmt.Fprint(w, "创建新的文章")
		}
	})

	http.ListenAndServe(":9191", router)

}
