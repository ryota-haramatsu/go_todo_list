## go_todo_list
### 内容
- DBの初期化処理と接続 (sqlite3)
- Userの作成とAuth機能
- Session処理
- Todo(CRUD)機能の作成

### メモ
- go module vendoringなど
go modules
https://tip.golang.org/doc/go1.16#tools

#### 初期化
1. ディレクトリを作成
2. go fmtをしてみたが怒られる
```
go: inconsistent vendoring in /Users/haramatsu:
	golang.org/x/tools@v0.1.0: is explicitly required in go.mod, but not marked as explicit in vendor/modules.txt

	To ignore the vendor directory, use -mod=readonly or -mod=mod.
	To sync the vendor directory, run:
		go mod vendor

```

3. どうやら初期化が必要らしいので
go mod init example.com/go_todoapp (example.com/はいらないかも)
example.com/go_todoappは引数としてモジュール名を指定する
https://qiita.com/uchiko/items/64fb3020dd64cf211d4e

すると
```
go: creating new go.mod: module example.com/go_todoapp
go: to add module requirements and sums:
	go mod tidy
```
と表示され、go.modファイルが作成される

iniが自動importされなかったが
```
go get "gopkg.in/go-ini/ini.v1"
```
してなかっただけ

ioパッケージ
https://christina04.hatenablog.com/entry/golang-io-package-diagrams


こんな感じのmain.goを作成したが、
```
func main() {
	log.Println("test")
	fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
}
```

なぜか下記のようなエラー
```
2021/04/27 23:39:35 open : no such file or directory
exit status 1
```
config.go内のConfig変数の値が間違っていただけ
→成功するとwebapp.logファイルが出力されてログが作成されていることがわかる

#### DB作成
①　DBの作成 + Userテーブルの設定、作成
sqlite3のインストール
- go get github.com/mattn/go-sqlite3
- sqlite3を使用しているので　webapp.sqlを作成したら 「sqlite3 webapp.sql」
- .table でテーブルを確認
- usersテーブルにuserを作成するコマンドを作成
- uuid パッケージをgetしておく
https://github.com/google/uuid

uuid = いつでも誰でも作れるけど、作ったIDは世界中で重複しないことになっているID
https://wa3.i-3-i.info/word13163.html

DB Open
https://golang.org/pkg/database/sql/#Open

```
The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections. 
Thus, the Open function should be called just once. 
It is rarely necessary to close a DB.
```
- sql.Open()は接続を確立するのではなく、抽象化した構造体を返す
- dbはConnection Poolを利用するため一度Open()したものをClose()せずに使いまわすことが基本
- Open(), Close()を頻繁にすると利用効率の低下、ネットワーク帯域の圧迫、TCPのTIME_WAITなどが発生する?

結果セットの取得
Query() と Exec()
- Query(): 複数の検索結果を取得する場合に使用する。1行の場合はQueryRow()
- Exec(): INSERT UPDATEなど検索結果を取得しない場合に使用する

②　ユーザーの作成（Create）
Userを作成しようとしたとき
0001-01-01 00:00:00 +0000 UTC
という表記になっておりエラー

Panic
```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x20 pc=0x40eb921]
```
宣言したerr への再代入を = ではなく := として入れてしまったから

③　ユーザーの取得（Read）
func GetUser()
コマンドを作成
Db.QueryRow().Scan()を実行
④　ユーザーの更新（Update）
⑤　ユーザーの削除（Delete) 

