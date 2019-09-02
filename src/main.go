package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	logger *log.Logger
)

var rootDirPath = os.Getenv("GOPATH") + os.Getenv("PROJECT_PATH")

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/ping", pingServiceHandler)

	return r
}

func main() {
	setLogFiles()
	r := setupRouter()
	r.Run(":8080")
}

func setLogFiles() {
	if _, err := os.Stat(rootDirPath + "/log"); os.IsNotExist(err) {
		os.Mkdir(rootDirPath+"/log", os.ModePerm)
	}
	setGinLog()
	setErrorLog()
}

func setGinLog() {
	gin.DisableConsoleColor()
	f, _ := os.Create(rootDirPath + "/log/gin_info.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func setErrorLog() {
	f, err := os.OpenFile(rootDirPath+"/log/error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger = log.New(f, "API  ", log.LstdFlags)
	logger.Println("Error log starts")
}
