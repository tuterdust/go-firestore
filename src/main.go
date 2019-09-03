package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	logger      *log.Logger
	rootDirPath string
)

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
	setPath()
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

func setPath() {
	rootDirPath = getEnv("GOPATH", "$HOME/go") + getEnv("PROJECT_PATH", "/src/github.com/tuterdust/go-firestore")
}

func getEnv(key, defaultValue string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	return defaultValue
}
