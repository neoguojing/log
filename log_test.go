package log

import "testing"

func TestLog(t *testing.T) {
	Info("hello world")
	log := NewLogger()
	log.Info("NewLogger")
}
