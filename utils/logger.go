package utils

import (
	"fmt"
	"os"
	"time"
)

var level_cnf = map[string]string{
	"DEBUG": "95",
	"INFO":  "96",
	"WARN":  "33",
	"ERROR": "31",
}

func ErrorLog(s string, a ...interface{}) {
	if len(a) == 0 {
		log("ERROR", s)
	} else {
		log("ERROR", s, a)
	}
}

func WarnLog(s string, a ...interface{}) {
	if len(a) == 0 {
		log("WARN", s)
	} else {
		log("WARN", s, a)
	}
}

func InfoLog(s string, a ...interface{}) {
	if len(a) == 0 {
		log("INFO", s)
	} else {
		log("INFO", s, a)
	}
}

func DebugLog(s string, a ...interface{}) {
	if os.Getenv("DEBUG") != "true" {
		return
	}

	if len(a) == 0 {
		log("DEBUG", s)
	} else {
		log("DEBUG", s, a)
	}
}

func log(level string, s string, a ...interface{}) {
	currentTime := time.Now().Format("[2006-01-02 15:04:05]")

	if len(a) == 0 {
		fmt.Printf(currentTime+" \033[1m\033[%sm[%s] %s\033[0m\n", level_cnf[level], level, s)
	} else {
		fmt.Printf(currentTime+" \033[1m\033[%sm[%s] %s: %s\033[0m\n", level_cnf[level], level, s, a[0])
	}
}
