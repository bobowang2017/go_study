package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	// 首先创建一个函数字典用于注册函数
	funcMap := template.FuncMap{
		// 注册函数title, strings.Title会将单词首字母大写
		"title": strings.Title,
	}
	const templateText = `
		 Input: {{printf "%q" .}}
		 Output 0: {{title .}}
		 Output 1: {{title . | printf "%q"}}
		 Output 2: {{printf "%q" . | title}}
	`

	// 创建模板, 添加模板函数,添加解析模板文本.
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// 运行模板，出入数据参数
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}
