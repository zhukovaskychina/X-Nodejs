package main

import (
	"fmt"
	"github.com/zhouhui8915/go-socket.io-client"
	"log"
	"os"
)

var client *socketio_client.Client

func GetSocketIOInstance() (*socketio_client.Client, error) {
	if client == nil {
		log.Printf("start to connect the remote server\n")
		opts := &socketio_client.Options{
			Transport: "websocket",
			Query:     make(map[string]string),
		}
		config, err := ReadConfig("./NodeAgentConf.conf")
		APPID := config.AppId
		IP := config.LocalMachineIpAddress
		REMOTEURL := config.RemoteServer

		if APPID == "" {

			PrintErrorMsg("APPID 需要在环境变量中进行设置!")
			os.Exit(3)
		}
		if IP == "" {

			PrintErrorMsg("IP 需要在环境变量中进行设置!")
			os.Exit(3)
		}
		if REMOTEURL == "" {

			PrintErrorMsg("服务器端地址REMOTEURL 需要在环境变量中进行设置!")
			os.Exit(3)
		}
		opts.Query["APPID"] = APPID
		opts.Query["IP"] = IP

		uri := REMOTEURL

		clients, err := socketio_client.NewClient(uri, opts)
		if err != nil {
			log.Printf("NewClient error :%v\n", err)
			return nil, err
		}
		if clients.GetConnectState() == 2 {
			log.Printf("Mission complete!")
		}

		client = clients
		clients = nil
		var sendMap map[string]string
		sendMap = make(map[string]string)
		client.Emit("connectSuccess", sendMap)
	} else {

		if (client.GetConnectState() != 1) && (client.GetConnectState() != 2) {

			log.Printf("the agent will be delayed in 10 minutes")
			//	time.Sleep(600000*time.Second)
			log.Printf("restart to connect the remote server\n")
			client.Reconnect()
			fmt.Println(client.GetConnectId())
			if client.GetConnectState() == 2 || client.GetConnectState() == 1 {
				var sendMap map[string]string
				sendMap = make(map[string]string)
				client.Emit("reconnectSuccess", sendMap)
			}
			fmt.Println(client.GetConnectState())
			log.Printf("restart complete!\n")
		}
	}

	return client, nil
}

func PrintErrorMsg(errorMsg string) {
	fmt.Println("======================================NodejsAgent 无法正常启动!==============================================")
	fmt.Println("	异常原因:" + errorMsg)
	fmt.Println("============================================================================================================")
}
