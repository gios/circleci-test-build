package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	router "github.com/gios/mom-and-i/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	logFolder      = "logs"
	logTimePattern = "2006_01_02"
)

func init() {
	initLogger()
	initEnvLoad()
}

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	if _, err := os.Stat(logFolder); os.IsNotExist(err) {
		err := os.Mkdir(logFolder, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	currentTime := time.Now().Local()
	logFilePath := fmt.Sprintf("./%s/%s.log", logFolder, currentTime.Format(logTimePattern))
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	log.SetLevel(log.WarnLevel)
}

func initEnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Warn(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	r := router.InitRoutes()
	http.Handle("/", r)
	fmt.Println("Application running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
