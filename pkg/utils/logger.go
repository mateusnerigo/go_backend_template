package utils

import (
	"log"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

type Logger struct {
	FileName string
}

func NewLogger(fileName string) *Logger {
	_, cFileName, _, ok := runtime.Caller(1)

	if fileName == "" && ok {
		fileName = cFileName
	}

	return &Logger{FileName: fileName}
}

func (l *Logger) Log(message any) {
	godotenv.Load()
	debug := os.Getenv("DEBUG")

	if debug == "TRUE" {
		log.Printf("\033[34mFILE: \033[0m%v", l.FileName)
		log.Printf("\033[32mLOG: \033[0m%v", message)
	}
}

func (l *Logger) Warn(message any) {
	godotenv.Load()
	debug := os.Getenv("DEBUG")

	if debug == "TRUE" {
		log.Printf("\033[34mFILE: \033[0m%v", l.FileName)
		log.Printf("\033[33mWARN: \033[0m%v", message)
	}
}

func (l *Logger) Error(message any) {
	godotenv.Load()
	debug := os.Getenv("DEBUG")

	if debug == "TRUE" {
		log.Printf("\033[34mFILE: \033[0m%v", l.FileName)
		log.Printf("\033[31mERROR: \033[0m%v", message)
	}
}
