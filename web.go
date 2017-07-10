package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

type TemplateData struct {
	Title   string
	Source  string
	Execute string
	Define  string
}

var source string
var execute string
var define_slice []string
var define_func_slice []string

func HelloServer(w http.ResponseWriter, r *http.Request) {
	// 親テンプレートindex.htmlと、小テンプレートbody.htmlを用意
	tmpl := template.Must(template.ParseFiles("views/index.html", "views/body.html"))

	// タイトル
	title := "Racket実行環境"

	// ソースコード
	source_html := strings.Replace(source, "\n", "<br>", -1)

	// 実行結果
	execute_html := strings.Replace(execute, "\n", "<br>", -1)

	// define で定義された関数を抜き出す
	var define_html string
	for _, arg := range define_slice {
		define_html += "<option value=\"" + arg + "\">" + arg + "</option>"
	}

	// テンプレートを実行して出力
	templatedata := TemplateData{title, source_html, execute_html, define_html}
	if err := tmpl.ExecuteTemplate(w, "base", templatedata); err != nil {
		fmt.Println(err)
	}
}

// Execute されたら
func Post_exec(w http.ResponseWriter, r *http.Request) {
	// 初期化
	source = ""
	// define_slice を source に追加
	for _, arg := range define_slice {
		source += arg + "\n"
	}
	// 入力を source に追加
	source += r.FormValue("sourcecode")
	// debug
	fmt.Println("> receive source")
	fmt.Println(source)
	//ここに source > define を書く
	//TODO: \n区切りでsliceを切る > 式が改行されずに書かれているのにも対応
	slice := strings.Split(source, "\n")
	for _, arg := range slice {
		// "define" があるか
		if strings.HasPrefix(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(arg, " "), "("), " "), "define") {
			// func_name に関数名代入
			tmp := strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(arg, " "), "("), " "), "define"), " "), "("), " ")
			re := regexp.MustCompile(".*? ")
			func_name := re.FindString(tmp)
			// define_func_slice に func_name が登録されていないか
			flag := 0
			for _, arg2 := range define_func_slice {
				if func_name == arg2 {
					flag = 1
					break
				}
			}
			// もし登録されていなければ
			if flag == 0 {
				define_func_slice = append(define_func_slice, func_name)
				define_slice = append(define_slice, arg)
			}
		}
	}
	// 実行
	out, _ := exec.Command("racket", "-e", strings.Replace(source, "\n", " ", -1)).CombinedOutput()
	execute = string(out)
	// debug
	fmt.Println("< send result")
	fmt.Println(execute)
	// redirect
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", "/")
	w.WriteHeader(http.StatusMovedPermanently)
}

func Post_curry(w http.ResponseWriter, r *http.Request) {
	source = r.FormValue("definelist")
	fmt.Println(source)
	// currying して、重複チェックして、define_slice に append
	// TODO
	// redirect
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", "/")
	w.WriteHeader(http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/post_exec", Post_exec)
	http.HandleFunc("/post_curry", Post_curry)
	log.Printf("Start Go HTTP Server")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
