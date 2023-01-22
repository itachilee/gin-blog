package global

import (
	"github.com/itachilee/ginblog/pkg/setting"
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	AppSetting = &setting.App{}

	ServerSetting = &setting.Server{}

	DatabaseSetting = &setting.Database{}

	RedisSetting = &setting.Redis{}

	MqttSetting = &setting.Mqtt{}

	JWTSetting = &setting.JWT{}
)

func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini' :%v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
	err = Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
	err = Cfg.Section("mqtt").MapTo(MqttSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo MqttSetting err: %v", err)
	}

	err = Cfg.Section("jwt").MapTo(JWTSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo JwtSetting err: %v", err)
	}

}
