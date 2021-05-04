package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"example.com/go_todoapp/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	// filenamesをfilesのスライスに格納する処理
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// テンプレートをパース
	templates := template.Must(template.ParseFiles(files...))

	// defineでテンプレートを指定したとき、layoutを明示的に指定する必要がある
	templates.ExecuteTemplate(w, "layout", data)
}

//サーバーを立ち上げるためのコードを記述
func StartMainServer() error {
	// css, jsファイルを読み込む
	// http.FileServerで静的ページを返す
	files := http.FileServer(http.Dir(config.Config.Static))

	// static配下にcss,jsが設定されていないのでhttp.StripPrefixを使用してstaticを取り除く
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URLを追加していく
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)

	// 第二引数はnilだとページがなければ404エラーを返す
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
