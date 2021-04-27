package config

import (
	"log"

	"example.com/go_todoapp/utils"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

//外部から呼び出しできるようにグローバル化する
//ここでconfig.iniを読み込む
var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	//config.iniの[Section名]
	//X=YとなっているXがKey(X)
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("logfile").Key("logfile").String(),
	}
}
