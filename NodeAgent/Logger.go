package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var LOGGER *log.Logger

func mains() {
	fileName := "Info_First.log"
	logFile, err := os.Create(fileName)

	defer logFile.Close()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Fatalln("open file error")
	}
	debugLog := log.New(logFile, "[Info]", log.Llongfile)

	debugLog.Println("A Info message here")
	debugLog.SetPrefix("[Debug]")
	debugLog.Println("A Debug Message here ")
}

func GetLoggerInstance() *log.Logger {
	if LOGGER == nil {
		fileName := ".NodeAgentLog.log"

		logFile, err := os.Create(fileName)
		if err != nil {

		}
		defer logFile.Close()
		logger := log.New(logFile, "[Info]", log.Llongfile)
		LOGGER = logger

	}
	return LOGGER
}

func LOGDEBUG(message string) {
	logger := GetLoggerInstance()
	logger.SetPrefix("[Debug]" + GetCurrentDateNow())
	logger.Println(message)
}
