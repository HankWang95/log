package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

const (
	info  = "[INFO] "
	debug = "[DEBUG]"
	erro  = "[ERROR]"
)

var (
	logFile *os.File
	tmpLog  *os.File
)

func NewLogger(name, logPath string) {
	var err error
	logFile, err = os.OpenFile(path.Join(logPath, name+".log"), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		writeTmpLog()
	}
}

func Error(v ...interface{}) {
	checkLogFile()
	line := getCodeLine(2)
	t := getTimeFormat()
	s := red(fmt.Sprintf("%s [%s] [%s] %s", erro, t, line, fmt.Sprintln(v...)))
	logFile.WriteString(s)
}

func Debug(v ...interface{}) {
	checkLogFile()
	line := getCodeLine(2)
	t := getTimeFormat()
	s := yellow(fmt.Sprintf("%s [%s] [%s] %s", debug, t, line, fmt.Sprintln(v...)))
	logFile.WriteString(s)

}

func Info(v ...interface{}) {
	checkLogFile()
	t := getTimeFormat()
	s := green(fmt.Sprintf("%s [%s] %s", info, t, fmt.Sprintln(v...)))
	logFile.WriteString(s)

}

func getCodeLine(depth int) (content string) {
	_, file, line, ok := runtime.Caller(depth)
	//f := runtime.FuncForPC(pc)
	_, f := path.Split(file)
	if ok {
		//content = fmt.Sprint(f.Name(), ":", line)
		content = fmt.Sprint(f, ":", line)
	}
	return content
}

func getTimeFormat() (fmtTime string) {
	now := time.Now()
	fmtTime = now.Format("2006-01-02 15:04:05")
	return fmtTime
}

func checkLogFile() {
	if logFile == nil {
		writeTmpLog()
	}
}

func writeTmpLog() {
	line := getCodeLine(3)
	t := getTimeFormat()
	warning := fmt.Sprint("ERROR! ", t, " logfile not init ", line)
	fmt.Println(warning)
}
