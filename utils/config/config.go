// config 配置文件相关结构定义
package config

import (
    "github.com/go-ini/ini"
    "log"
    "time"
)

// Server Gin使用到的相关配置
type Server struct {
    RunMode      string
    HttpPort     int
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
    LogDir       string
}

var ServerSetting = &Server{}


// Redis 相关配置
type Redis struct {
    Host               string
    Password           string
    MaxIdle            int
    MaxActive          int
    IdleTimeout        time.Duration
    SelectDB           int
    DialConnectTimeout time.Duration
    DialReadTimeout    time.Duration
    DialWriteTimeout   time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File
var CFGFilePath string

func Setup() {
    var err error
    cfg, err = ini.Load("app.ini")
    if err != nil {
        log.Fatalf("setting.Setup, fail to parse 'app.ini': %v", err)
    }

    mapTo("server", ServerSetting)
    mapTo("redis", RedisSetting)

    ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
    ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}

func mapTo(section string, v interface{}) {
    err := cfg.Section(section).MapTo(v)
    if err != nil {
        log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
    }
}
