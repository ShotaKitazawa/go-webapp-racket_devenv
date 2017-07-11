package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"text/template"
)

type TemplateData struct {
	Title  string
	Source string
	Result string
	Define string
}

var source string
var result string
var define_slice []string
var define_func_slice []string

func HelloServer(w http.ResponseWriter, r *http.Request) {
	// 親テンプレートindex.htmlと、小テンプレートbody.htmlを用意
	tmpl := template.Must(template.ParseFiles("views/index.html", "views/body.html"))

	// タイトル
	title := "Racket実行環境"

	// ソースコード
	//source_html := strings.Replace(source, "\n", "<br>", -1)
	source_html := strings.Replace(source, "\n", "<br>", -1)

	// 実行結果
	result_html := strings.Replace(result, "\n", "<br>", -1)

	// define で定義された関数を抜き出す
	var define_html string
	for _, arg := range define_slice {
		define_html += "<option value=\"" + arg + "\">" + arg + "</option>"
	}

	// テンプレートを実行して出力
	templatedata := TemplateData{title, source_html, result_html, define_html}
	if err := tmpl.ExecuteTemplate(w, "base", templatedata); err != nil {
		fmt.Println(err)
	}
}

// Execute されたら
func Post_exec(w http.ResponseWriter, r *http.Request) {
	// 初期化
	source = ""
	result = ""
	// define_slice を source に追加
	for _, arg := range define_slice {
		source += arg + "\n"
	}
	// 入力を source に追加
	// TODO: sourcecode からの入力をまずフォーマット
	//reContinuousBlank := regexp.MustCompile(" +")
	//source += reContinuousBlank.ReplaceAllString(r.FormValue("sourcecode"), " ")
	source += r.FormValue("sourcecode")
	// debug
	fmt.Println("> receive source")
	fmt.Println(source)
	// 実行
	out, err := exec.Command("racket", "-e", strings.Replace(source, "\n", " ", -1)).CombinedOutput()
	// source > defined
	if err == nil {
		add_defined(source)
	}
	// 実行結果代入
	result = string(out)
	// debug
	fmt.Println("< send result")
	fmt.Println(result)
	// redirect
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", "/")
	w.WriteHeader(http.StatusMovedPermanently)
}

func Post_curry(w http.ResponseWriter, r *http.Request) {
	// 初期化
	source = ""
	result = ""
	// 入力を source に追加
	source = strings.Replace(r.FormValue("definelist"), "\n", "", -1)
	fmt.Println("> receive source")
	fmt.Println(source)
	submit := r.FormValue("submit")
	if submit == "Curry" {
		// Curry化: header.go
		result = currying(source)
		clearing(source)
		add_defined(result)
	} else if submit == "Clear" {
		clearing(source)
		result = "Delete element"
	}
	fmt.Println("< send result")
	fmt.Println(result)
	// redirect
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", "/")
	w.WriteHeader(http.StatusMovedPermanently)
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/post_exec", Post_exec)
	http.HandleFunc("/post_curry", Post_curry)
	log.Printf("Start Go HTTP Server")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
