package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

var router = mux.NewRouter().StrictSlash(true)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("%s", "<h1>Hello, 欢迎来到 goblog！</h1>"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "这是一个blog示例，有问题请联系："+"<a href='mailto:518@msn.cn'>518@msn.cn</a>")
}
func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	path := vars["path"]
	fmt.Fprint(w, "文章 ID："+id+""+path+"xxx")
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")
	errors := make(map[string]string)
	//验证title内容
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if len(title) < 3 || len(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}
	//验证body内容
	if body == "" {
		errors["body"] = "内容不能为空"
	} else if len(body) < 10 {
		errors["body"] = "内容长度需大于或等于 10 个字节"
	}
	// 检查是否有错误
	if len(errors) == 0 {
		fmt.Fprint(w, "验证通过!<br>")
		fmt.Fprintf(w, "title 的值为: %v <br>", title)
		fmt.Fprintf(w, "title 的长度为: %v <br>", len(title))
		fmt.Fprintf(w, "body 的值为: %v <br>", body)
		fmt.Fprintf(w, "body 的长度为: %v <br>", len(body))
	} else {
		html := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<title>创建文章 —— 我的技术博客</title>
				<style type="text/css">.error {color: red;}</style>
			</head>
			<body>
				<form action="{{ .URL }}" method="post">
					<p><input type="text" name="title" value="{{ .Title }}"></p>
					{{ with .Errors.title }}
					<p class="error">{{ . }}</p>
					{{ end }}
					<p><textarea name="body" cols="30" rows="10">{{ .Body }}</textarea></p>
					{{ with .Errors.body }}
					<p class="error">{{ . }}</p>
					{{ end }}
					<p><button type="submit">提交</button></p>
				</form>
			</body>
			</html>
			`
		storeURL, _ := router.Get("articles.store").URL()
		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		}
		templ, err := template.New("create-form").Parse(html)
		if err != nil {
			panic(err)
		}
		if err = templ.Execute(w, data); err != nil {
			panic(err)
		}

	}
}

func forceHTMLMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
		h.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		// 2. 将请求传递下去
		next.ServeHTTP(w, r)
	})
}

func articleCreateHandler(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<title>创建文章 —— 我的技术博客</title>
		</head>
		<body>
			<form action="%s?test=data" method="post">
				<p><input type="text" name="title"></p>
				<p><textarea name="body" cols="30" rows="10"></textarea></p>
				<p><button type="submit">提交</button></p>
			</form>
		</body>
		</html>
		`
	storeURL, _ := router.Get("articles.store").URL()
	fmt.Fprintf(w, html, storeURL)
}

func main() {

	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")
	router.HandleFunc("/articles/{id:[0-9]+}/{path:.*}", articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles/create", articleCreateHandler).Methods("GET").Name("articles.create")

	// 自定义 404 页面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	// 中间件：强制内容类型为 HTML
	router.Use(forceHTMLMiddleware)
	// 通过命名路由获取 URL 示例
	/*
		homeURL, _ := router.Get("home").URL()
		fmt.Println("homeURL: ", homeURL)
		articleURL, _ := router.Get("articles.show").URL("id", "23", "path", "aaa")
		fmt.Println("articleURL: ", articleURL)
	*/
	http.ListenAndServe(":9191", removeTrailingSlash(router))

}
