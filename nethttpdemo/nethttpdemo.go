package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func SayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("访问的URL是: ", r.URL)
	fmt.Println("访问URL的方法是: ", r.Method)
	fmt.Println("访问的协议是: ", r.Proto)
	fmt.Println("Header信息为: ", r.Header)
	fmt.Fprint(w, "Keyman")
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			log.Fatal("模板报错 ", err)
		}
		t.Execute(w, nil)
	} else {
		//r.ParseForm()
		//fmt.Println(r.Form)
		username := template.HTMLEscapeString(r.FormValue("username"))

		fmt.Println(username)
		fmt.Fprint(w, username)
	}

	//log.Println()
}

func main() {
	http.HandleFunc("/sayhello", SayHelloName)
	http.HandleFunc("/login", Login)
	err := http.ListenAndServe(":9191", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
