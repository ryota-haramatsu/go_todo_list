package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// ログファイルの読み込み 読み書き、作成追記を指定
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("%v\n%T", err, err)
	}

	multiLogFile := io.MultiWriter(os.Stdout, logfile)   //出力先を標準出力とログファイル2つ変数に格納
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //フォーマット指定
	log.SetOutput(multiLogFile)                          //ログの出力先を指定
}
