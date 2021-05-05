package controllers

import (
	"log"
	"net/http"

	"example.com/go_todoapp/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)

	// 一致するセッションがなければ
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		// ユーザーを取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		// ユーザーが所持するTodoを取得
		todos, _ := user.GetTodosByUser()
		user.Todos = todos

		// 第二引数にuserを渡す
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

// todoの新規作成ページ
func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

// todo新規作成処理
func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm() // フォームの解析
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

// todoの編集ページ
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		// ユーザーの確認
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		// 引数idからtodoを取得
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

// todo編集機能
func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm() // フォームの解析
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession() // ユーザー取得
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		t := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}

// todo 削除機能
func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Fatalln(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}
		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}
