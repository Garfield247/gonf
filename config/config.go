package config

import (
	"time"

	"gopkg.in/ini.v1"
)

type SqlDataBase struct {
	Path string
}

type Jwt struct {
	SecterKey  string
	LifeLength time.Duration
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Log struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type Redis struct {
	Host     string
	Port     int
	Password string
	DB       int
}

var (
	DataBaseCfg = &SqlDataBase{}
	JwtCfg      = &Jwt{}
	ServerCfg   = &Server{}
	LogCfg      = &Log{}
	RedisCfg    = &Redis{}
)

func SetUp() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		panic(err)
	}
	if err := cfg.Section("sqlite").MapTo(DataBaseCfg); err != nil {
		panic(err)
	}
	if err := cfg.Section("jwt").MapTo(JwtCfg); err != nil {
		panic(err)
	}
	if err := cfg.Section("server").MapTo(ServerCfg); err != nil {
		panic(err)
	}
	if err := cfg.Section("log").MapTo(LogCfg); err != nil {
		panic(err)
	}
	if err := cfg.Section("redis").MapTo(RedisCfg); err != nil {
		panic(err)
	}
}
