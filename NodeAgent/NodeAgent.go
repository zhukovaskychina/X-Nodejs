package main

import (
	"fmt"
	"github.com/zhouhui8915/go-socket.io-client"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type SystemInfo struct {
	loadAvg    Loadavg
	memStatus  MemStatus
	diskStatus DiskStatus
}

var currentNum int = 0
var pidMap map[string]string

var socketioclient socketio_client.Client

func main() {

	pidMap = make(map[string]string)

	nodeAgentCount := getNodeAgentCount()
	if nodeAgentCount >= 2 {
		PrintErrorMsg("已经有了一个NodeAgent进程启动!请勿重复启动进程!")
		os.Exit(3)
	}
	fmt.Println("============================================================================================================")
	fmt.Println("					All Rights Reserved @zhukovasky zhukovasky开发监控神器									")
	fmt.Println("							如有问题请联系作者									 						    ")
	fmt.Println("使用说明:")
	fmt.Println("请在环境变量当中配置APPID，IP，REMOTEURL信息")
	fmt.Println("APPID:该node.js服务名称")
	fmt.Println("IP:该node.js 服务器IP地址")
	fmt.Println("REMOTEURL:远程服务器数据接收地址")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Agent设计说明:")
	fmt.Println("该agent的作用，将把定制node.js生成的日志，以每分钟的频率向远程服务器发送，并且将接收远程服务器端的指令，用于打dump")
	fmt.Println("用于上传dump文件，实现云端分析node.js服务快照实例")
	fmt.Println("============================================================================================================")
	NODEJS_LOG_DIR := os.Getenv("NODEJS_LOG_DIR")
	LOGDEBUG("该agent的作用，将把定制node.js生成的日志，以每分钟的频率向远程服务器发送，并且将接收远程服务器端的指令，用于打dump")
	if NODEJS_LOG_DIR == "" {
		fmt.Println("请在环境变量当中配置该变量值具体直!")
		os.Exit(3)
	}

	//NODEJS_ERRORLOG_DIR:=os.Getenv("NODEJS_ERROR_LOG_DIR")
	tmpLockPositionFile := NODEJS_LOG_DIR + "/" + ".tmpLockPositionFile"

	isExist, err := PathExists(tmpLockPositionFile)

	if err != nil {
		return
	}
	if !isExist {
		createFile(tmpLockPositionFile)
	}
	socketioclient, err := GetSocketIOInstance()
	if err != nil {
		return
	}
	w := New()
	w.SetMaxEvents(1)
	w.FilterOps(Write, Create)
	go func() {
		for {
			select {
			case event := <-w.Event:
				DateNow := Format("YYYY-MM-DD", time.Now())
				fileName := "PingAnNode" + DateNow + ".log"
				changeFileName := event.FileInfo.Name()
				if event.Op.String() == "CREATE" {
					//检测到有文件新生成
					if strings.HasSuffix(changeFileName, ".heapsnapshot") {
						
					}
					if strings.HasSuffix(changeFileName, ".heapprofile") {

					}
					if strings.HasSuffix(changeFileName,".cpuprofile") {
						
					}
				}
				if event.Op.String() == "WRITE" {
					if strings.Compare(changeFileName, fileName) == 0 {
						changeFilePath := NODEJS_LOG_DIR + "/" + changeFileName
						size := ReadLineForCountNum(changeFilePath)
						delta := size - currentNum
						if delta > 0 {
							ReadLineByPosition(changeFilePath, sendNodeJsInfoToRemoteServer, currentNum)
						}
						currentNum = size
						writeContentToSomeFile(tmpLockPositionFile, DateNow+";"+strconv.Itoa(currentNum)+"")
					} else {

					}
				}

			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	socketioclient.On("error", func() {
		LOGDEBUG("on error\n")
	})
	socketioclient.On("connection", func() {
		log.Printf("on connect\n")
		socketioclient.Emit("connectSuccess", nil)
	})
	socketioclient.On("message", func(msg string) {
		log.Printf("on message:%v\n", msg)
	})
	socketioclient.On("disconnection", func() {
		log.Printf("on disconnect\n")
		log.Printf("远程服务器停止服务!\n")
		os.Exit(3)
		socketioclients, errs := GetSocketIOInstance()
		if errs == nil {
			socketioclient = socketioclients
		}
		log.Printf("reconnect the remote server")
	})

	socketioclient.On("heapdump", func(msg string) {
		command := "NodeKiller " + msg + " --heapdump"
		exeCommand(command)
	})
	socketioclient.On("cpuProfile", func(msg string) {
		command := "NodeKiller " + msg + " --cpuProfile"
		exeCommand(command)
	})
	socketioclient.On("heapProfile", func(msg string) {
		command := "NodeKiller " + msg + " --heapProfile"
		exeCommand(command)
	})
	socketioclient.On("traceGCVerboseNvp", func(msg string) {
		command := "NodeKiller " + msg + " --traceGCVerboseNvp"
		exeCommand(command)
	})
	socketioclient.On("forceGC", func(msg string) {
		command := "NodeKiller " + msg + " --forceGC"
		exeCommand(command)
	})

	socketioclient.On("uploadDumpFile", func(msg string) {

	})
	// Watch this folder for changes.
	if err := w.Add(os.Getenv("NODEJS_LOG_DIR")); err != nil {
		log.Fatalln(err)
	}

	// Watch test_folder recursively for changes.
	if err := w.AddRecursive(os.Getenv("NODEJS_LOG_DIR")); err != nil {
		log.Fatalln(err)
	}

	go func() {
		w.Wait()
		w.TriggerEvent(Create, nil)
		w.TriggerEvent(Remove, nil)
	}()

	go func() {
		timer(timerFunc)
	}()
	go func() {
		systemTimer(sendSystemInfoToRemoteServer)
	}()
	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 60000); err != nil {
		log.Fatalln(err)
	}
}

func timerFunc() {
	//说明有新的进程ID 加入进来

	//socketioclient.Emit("ProcessEventExitOrJoinIn", sendMap)
	//获取当前node.js进程信息
	pids := getNodejsProcessPid()
	evenType, sendPids := calculateMissPid(pidMap, convertArrayToMap(pids))
	if evenType == "UPDATEDOWNPID" {
		//如果有进程退出，必然有进程启动
		sendAddPids := calculateAddPid(pidMap, convertArrayToMap(pids))
		var sendMap map[string][]string
		sendMap = make(map[string][]string)
		sendMap["AddPids"] = sendAddPids
		sendMap["DownPids"] = sendPids
		socketioclient, err := GetSocketIOInstance()
		if err != nil {
			return
		}
		socketioclient.Emit("ProcessEventExitOrJoinIn", sendMap)
		var sendProcessMap map[string]string
		sendProcessMap = make(map[string]string)
		for _, s := range sendAddPids {
			times := getProcessStartTime(s)
			sendProcessMap["startTime"] = times
			sendProcessMap["pid"] = s
			socketioclient.Emit("ProcessRestartEventJoinIn", sendProcessMap)
		}

	} else {
		//如果不是，则是进程退出，或者进程重新启动，此时需要将新增的进程ID 信息的启动时间发送给服务器
		var sendMap map[string][]string
		sendMap = make(map[string][]string)

		if evenType == "ADD" {

			sendMap["AddPids"] = sendPids
			socketioclient, err := GetSocketIOInstance()
			if err != nil {
				return
			}
			//进程启动信息
			socketioclient.Emit("ProcessEventJoinIn", sendMap)
			var sendProcessMap map[string]string
			sendProcessMap = make(map[string]string)
			for _, s := range sendPids {
				times := getProcessStartTime(s)
				sendProcessMap["startTime"] = times
				sendProcessMap["pid"] = s
				socketioclient.Emit("ProcessAddEventJoinIn", sendProcessMap)
			}
		} else {
			socketioclient, err := GetSocketIOInstance()
			if err != nil {
				return
			}
			sendMap["DownPids"] = sendPids
			socketioclient.Emit("ProcessEventExit", sendMap)
		}
	}
	//重置
	pidMap = convertArrayToMap(pids)
}

func timer(timer func()) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			timer()
		}
	}
}

func systemTimer(timer func()) {
	ticker := time.NewTicker(6 * time.Second)
	for {
		select {
		case <-ticker.C:
			timer()
		}
	}
}
func sendNodeJsInfoToRemoteServer(line string) {
	go func() {
		strArray := strings.Split(line, " ")
		timestamp := (strArray[0] + " " + strArray[1])
		level := (strArray[2])
		types := (strArray[3])
		pid := (strArray[4])
		content := (strArray[5])
		var sendMap map[string]string
		sendMap = make(map[string]string)
		sendMap["type"] = types
		sendMap["pid"] = pid
		sendMap["content"] = content
		sendMap["level"] = level
		sendMap["timestamp"] = timestamp
		client, err := GetSocketIOInstance()
		if err != nil {
			log.Printf("SocketClient error exist:%v\n", err)
			return
		}
		client.Emit("sendNodeInfo", sendMap)
	}()

}

func sendLoadAvgToRemoteServer() {
	go func() {
		var sendMap map[string]Loadavg
		sendMap = make(map[string]Loadavg)

		loadAvg, err := ParseLoadAvg()
		if err != nil {

		}
		client, err := GetSocketIOInstance()
		if err != nil {
			log.Printf("SocketClient error system:%v\n", err)
			return
		}
		sendMap["loadAvg"] = *loadAvg
		client.Emit("sendLoadAvg", sendMap)
	}()
}
func sendMemStatsToRemoteServer() {
	go func() {
		var sendMap map[string]MemStatus
		sendMap = make(map[string]MemStatus)

		client, err := GetSocketIOInstance()
		if err != nil {
			log.Printf("SocketClient error system:%v\n", err)
			return
		}
		memStatus := MemStat()
		sendMap["memStatus"] = memStatus
		client.Emit("sendMemStatus", sendMap)
	}()
}

func sendDiskStatusToRemoteServer() {
	go func() {
		var sendMap map[string]DiskStatus
		sendMap = make(map[string]DiskStatus)

		client, err := GetSocketIOInstance()
		if err != nil {
			log.Printf("SocketClient error system:%v\n", err)
			return
		}
		memStatus := DiskUsage("/")
		sendMap["diskStatus"] = memStatus
		client.Emit("sendDiskStatus", sendMap)
	}()
}

func sendSystemInfoToRemoteServer() {
	go func() {
		sendDiskStatusToRemoteServer()
		sendMemStatsToRemoteServer()
		sendLoadAvgToRemoteServer()
	}()

}
