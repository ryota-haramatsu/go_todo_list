package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

//ユーザーの作成を行う関数 *User型に対してメソッドを定義 ポインタレシーバ
func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) VALUES (?,?,?,?,?)`

	// DbやEncrypt(),createUUIDはmodelsパッケージに存在しているのでパッケージを指定しなくても使用できる
	_, err = Db.Exec(cmd,
		createUUID(), //uuid
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//ユーザーの取得機能
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`

	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

//ユーザーの更新機能
func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//ユーザーの削除機能
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

////////// セッション
// フォームに入力されたメールアドレスからユーザーを取得
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users
	where email = ?`
	//一つのレコードを取得したいのでQueryRow()
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

// セッション作成と取得処理
func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	// セッション作成
	cmd1 := `insert into sessions (
		uuid, 
		email, 
		user_id, 
		created_at) values (?,?,?,?)`
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	// ユーザーのセッション取得処理
	cmd2 := `select id, uuid, email, user_id, created_at
		from sessions where user_id = ? and email = ?`
	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
	)
	return session, err
}

// セッションがDBに存在するかどうか確認
func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at
		from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt,
	)
	// 存在しなければvalidをfalseにしてreturn
	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}
