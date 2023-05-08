package stanlog

import (
	"fmt"
	"time"
)

var logLevel int = 2
var logLevelMap map[string]int

func SetLogLevel(requestedLogLevel string) {
	logLevelMap := make(map[string]int)
	logLevelMap["DEBUG"] = 0
	logLevelMap["INFO"] = 1
	logLevelMap["WARNING"] = 2
	logLevelMap["ERROR"] = 3
	logLevelMap["CRITICAL"] = 4

	logLevel = logLevelMap[requestedLogLevel]
}

func SetLogLevelDebug() {
	logLevel = 0
}

func SetLogLevelInfo() {
	logLevel = 1
}

func SetLogLevelWarning() {
	logLevel = 2
}

func SetLogLevelError() {
	logLevel = 3
}

func SetLogLevelCritical() {
	logLevel = 4
}

func Debug(messageToLog string) {
	if logLevel == 0 {
		fmt.Printf("%s - %s    - %s\n", time.Now().Format("2006-01-02 15:04:05.000000"), "DEBUG", messageToLog)
	}
}

func Info(messageToLog string) {
	if logLevel < 2 {
		fmt.Printf("%s - %s     - %s\n", time.Now().Format("2006-01-02 15:04:05.000000"), "INFO", messageToLog)
	}
}

func Warning(messageToLog string) {
	if logLevel < 3 {
		fmt.Printf("%s - %s  - %s\n", time.Now().Format("2006-01-02 15:04:05.000000"), "WARNING", messageToLog)
	}
}

func Error(messageToLog string) {
	if logLevel < 4 {
		fmt.Printf("%s - %s    - %s\n", time.Now().Format("2006-01-02 15:04:05.000000"), "ERROR", messageToLog)
	}
}

func Critical(messageToLog string) {
	if logLevel < 5 {
		fmt.Printf("%s - %s - %s\n", time.Now().Format("2006-01-02 15:04:05.000000"), "CRITICAL", messageToLog)
	}
}
