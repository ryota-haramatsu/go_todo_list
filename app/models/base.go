package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"example.com/go_todoapp/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// テーブルのコードを作成

var Db *sql.DB

var err error

//テーブル名を宣言
const (
	tableNameUser = "users"
)

func init() {
	Db, err := sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	//コマンド
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at STRING)`, tableNameUser)
	Db.Exec(cmdU)
}

// uuid作成関数
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// パスワードをsha1でハッシュ化する関数
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
