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
