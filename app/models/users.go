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

//ユーザーの作成を行う関数 *User型に対してメソッドを定義 ポインタレシーバ
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?,?,?,?,?)`

	// DbやEncrypt(),createUUIDはmodelsパッケージに存在しているのでパッケージを指定しなくても使用できる
	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
