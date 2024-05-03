package logging

import (
	"log"
	"os"
)

var logFlag int
var InfoLog, ErrLog *log.Logger

func Logging() {
	logFlag = log.LstdFlags | log.Lshortfile
	InfoLog = log.New(os.Stdout, "INFO: ", logFlag)
	ErrLog = log.New(os.Stdout, "ERROR: ", logFlag)
}
