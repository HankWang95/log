package log

import (
	"errors"
	"testing"
)

var (
	hi = errors.New("new error")
)

func TestInitLogger(t *testing.T) {
	//NewLogger("testLog", "/Users/hank-for-work/Desktop/go/src/github.com/HankWang95/log")
}

func TestError(t *testing.T) {
	Error("test1", "are", "you", "ok?", hi)
}

func TestDebug(t *testing.T) {
	Debug("test2", "are", "you", "ok?")
}

func TestInfo(t *testing.T) {
	Info("test3", "are", "you", "ok?")
}
