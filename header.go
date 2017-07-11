package main

import (
	"fmt"
	"regexp"
	"strings"
)

func remove(words []string, search string) []string {
	result := []string{}
	for _, word := range words {
		if word != search {
			result = append(result, word)
		}
	}
	return result
}

func add_defined(source string) {
	//TODO: \n区切りでsliceを切る > 式が改行されずに書かれているのにも対応
	slice := strings.Split(source, "\n")
	for _, arg := range slice {
		// "define" があるか
		if strings.HasPrefix(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(arg, " "), "("), " "), "define") {
			// funcName に関数名代入
			tmp := strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(arg, " "), "("), " "), "define"), " "), "("), " ")
			reFuncName := regexp.MustCompile(".*? ")
			funcName := reFuncName.FindString(tmp)
			// define_func_slice に funcName が登録されていないか
			flag := 0
			for _, arg2 := range define_func_slice {
				if funcName == arg2 {
					flag = 1
					break
				}
			}
			// もし登録されていなければ
			if flag == 0 {
				define_func_slice = append(define_func_slice, funcName)
				define_slice = append(define_slice, arg)
			}
		}
	}
}

// currying 可能かチェックして、currying して、重複チェックして、define_slice に append
func currying(source string) string {
	// function = "f x" | "fc x y" | ...
	function := strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(source, " "), "("), " "), "define"), " "), "("), " ")
	// 途中で 関数名より左の文字列 を抜き出す
	reFuncName := regexp.MustCompile(".*? ")
	funcName := reFuncName.FindString(function)
	leftOfFunc := source[:strings.Index(source, funcName)]
	tmp := source[strings.Index(source, funcName):]
	rightOfFunc := strings.Replace(tmp[strings.Index(tmp, ")"):], "\n", "", -1)
	function = function[:strings.Index(function, ")")]
	reContinuousBlank := regexp.MustCompile(" +")
	function = strings.TrimRight(reContinuousBlank.ReplaceAllString(function, " "), " ")
	// 関数に二引数以上あれば Curring
	if strings.Index(function[strings.Index(function, " ")+1:], " ") != -1 {
		source = leftOfFunc + function[:strings.LastIndex(function, " ")] + ") (\\lambda (" + function[strings.LastIndex(function, " "):] + rightOfFunc + ")\n"
		fmt.Println(source)
		//再帰
		source = currying(source)

	}
	return source
}

func clearing(source string) {
	// define_slice, define_func_slice から該当要素の削除
	tmp := strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(source, " "), "("), " "), "define"), " "), "("), " ")
	reFuncName := regexp.MustCompile(".*? ")
	funcName := reFuncName.FindString(tmp)
	// 要素の削除: header.go
	define_slice = remove(define_slice, source)
	define_func_slice = remove(define_func_slice, funcName)
}
