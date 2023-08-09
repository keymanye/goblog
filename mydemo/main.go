package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var router = mux.NewRouter().StrictSlash(true)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 中间件
func forceHTMLMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
		h.ServeHTTP(w, r)
	})
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.gohtml")
	checkError(err)
	tmpl.Execute(w, nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "如果你在访问网站的过程中遇到任何问题请联系我们"+
		"<a href='mailto:518@msn.cn'>518@msn.cn</a>")
}

// 增加文章
func articleCreateHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/create.gohtml")
	templ.Execute(w, nil)
	checkError(err)
	//fmt.Fprint(w, "创建文章")
}

func articleStoreHandler(w http.ResponseWriter, r *http.Request) {
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
		lastInsertID, err := saveArticleToDB(title, body)
		if lastInsertID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatInt(lastInsertID, 10))
		} else {
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		storeURL, _ := router.Get("articles.store").URL()
		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		}


}

// 删除文档
func articleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "创建文章")
}

// 更新文章
func articleUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "创建文章")
}

// 编辑文章
func articlesEditHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "创建文章")
}

// 列出所有文档
func articleListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "创建文章")
}

func main() {
	router.Use(forceHTMLMiddleware)
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	router.HandleFunc("/home", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")
	router.HandleFunc("/articles/create", articleCreateHandler).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles/{id:[0-9]+}/edit", articlesEditHandler).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", articleUpdateHandler).Methods("POST").Name("articles.update")
	router.HandleFunc("/articles/{id:[0-9]+}/delete", articleDeleteHandler).Methods("POST").Name("articles.delete")
	http.ListenAndServe(":9090", router)
}
