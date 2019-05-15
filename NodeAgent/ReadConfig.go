package main

import (
	"gopkg.in/ini.v1"
	"log"
)

type Config struct {
	RemoteServer          string `ini:"remoteServer"`
	LocalMachineIpAddress string `ini:"localMachineIpAddress"`
	RemoteUploadUrl       string `ini:"remoteUploadUrl"`
	AppId                 string `ini:"appId"`
}

func ReadConfig(path string) (Config, error) {

	var config Config
	conf, err := ini.Load(path) //加载配置文件
	if err != nil {
		log.Println("加载配置文件失败!")
		return config, err
	}
	conf.BlockMode = false
	err = conf.MapTo(&config) //解析成结构体
	if err != nil {
		log.Println("配置文件可能格式不对!")
		return config, err
	}
	return config, err
}
