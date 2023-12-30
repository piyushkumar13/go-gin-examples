package logger

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"io"
	"logrusexample2/config"
	"os"
	"strings"
	"time"
)

func Init() {

	fmt.Println("Configuring Logger....")

	customFormatter := new(logger.JSONFormatter)
	customFormatter.TimestampFormat = time.RFC3339
	logger.SetFormatter(customFormatter)

	configDetails := config.GetConfig()

	loggingLevel := configDetails.GetString("logging.level")

	fmt.Println("Setting logging level", loggingLevel)
	setLogLevel(loggingLevel)
	logger.SetReportCaller(true) // it will add more details to the logs like function name and file name as well.

	//logrus.SetOutput(os.Stdout) // set the output to console

	file, _ := os.Create("logfile.log")
	logger.SetOutput(file) // sets the output to the file

	multiWriter := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(multiWriter) // sets the output to multiple destinations like file and console
}

func setLogLevel(level string) {
	switch strings.ToLower(level) {

	case "debug":
		logger.SetLevel(logger.DebugLevel)
	case "info":
		logger.SetLevel(logger.DebugLevel)
	case "error":
		logger.SetLevel(logger.DebugLevel)
	default:
		logger.SetLevel(logger.DebugLevel)
	}
}
