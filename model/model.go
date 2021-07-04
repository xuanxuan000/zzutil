package model

//BaseConfig 通用的服务配置
type BaseConfig struct {
	HTTPListenAddr string
	ServiceName    string
	BuildVersion   string // "default"

	Mysql map[string]MysqlCfg

	Log struct {
		LogLevel string
		Redirect bool
	}
}

type MysqlCfg struct {
	IP       string
	Port     string
	User     string
	PassWord string
	DBname   string
	SSL      bool //server是否要求ssl
}
