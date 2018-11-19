package util

import (
	"log"
	"os"
	"sync"
)

const loggerFilePath string = "/logs/externalauth.log"

var once sync.Once
var logger *log.Logger

func GetLogger() *log.Logger {
	once.Do(func() {
		logger = createLogger()
	})
	return logger
}

func createLogger() *log.Logger {
	file, _ := os.OpenFile(loggerFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	var rotateWritter = &RotateWriter{
		filename:loggerFilePath,
		file: file,
		lock: &sync.Mutex{},
	}
	return log.New(rotateWritter, "", log.LstdFlags)
}