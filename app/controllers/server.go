package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"text/template"

	"example.com/go_todoapp/app/models"
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

// アクセス制限用のセッションチェック
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	// requestからクッキーを取得
	cookie, err := r.Cookie("_cookie")
	// エラーがなければ
	if err == nil {
		sess = models.Session{UUID: cookie.Value} //DBにセッションがあるか確認
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/{id} という処理をしたい
		q := validPath.FindStringSubmatch(r.URL.Path) // マッチした部分をスライスで取得
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
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
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))

	// 第二引数はnilだとページがなければ404エラーを返す
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
