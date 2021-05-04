package controllers

import (
	"log"
	"net/http"

	"example.com/go_todoapp/app/models"
)

// Auth関係のハンドラを作成

// ユーザー新規登録処理ハンドラ
func signup(w http.ResponseWriter, r *http.Request) {
	// r.Methodでリクエストのステータスを取得
	if r.Method == "GET" {
		// signupページの出力
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm() // フォームの解析
		if err != nil {
			log.Println(err)
		}
		/* 入力の値(名前、メールアドレス、パスワード)を受けてとって
		ユーザーを作成したいのでユーザーのstructを作成する
		*/
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
		// w,r,リダイレクト先のURL、ステータスコードを指定し、リダイレクト
		http.Redirect(w, r, "/", 302)
	}
}

// ログイン処理ハンドラ
func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "login")
}

// ユーザー認証ハンドラ
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // フォームの解析
	if err != nil {
		log.Fatalln(err)
	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Fatalln(err)
		// エラーの場合はリダイレクト
		http.Redirect(w, r, "/login", 302)
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Fatalln(err)
		}
		// ここで作成したログイン情報からブラウザ側のcookieを作成
		// http.Cookie structに情報を入力
		// ここからCookie作成
		cookie := http.Cookie{
			Name:  "_cookie",
			Value: session.UUID,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
