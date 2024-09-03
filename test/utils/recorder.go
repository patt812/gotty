package utils

import (
	"bytes"
	"os"
	"strings"
	"sync"
)

type LogRecorder struct {
	oldStdout *os.File
	r         *os.File
	w         *os.File
	buffer    *bytes.Buffer
	mu        sync.Mutex
	started   bool
}

func NewLogRecorder() *LogRecorder {
	r, w, err := os.Pipe()
	if err != nil {
		panic("Failed to create pipe: " + err.Error())
	}

	oldStdout := os.Stdout
	os.Stdout = w

	return &LogRecorder{
		oldStdout: oldStdout,
		r:         r,
		w:         w,
		buffer:    &bytes.Buffer{},
		started:   true,
	}
}

func (lr *LogRecorder) stop() {
	lr.mu.Lock()
	defer lr.mu.Unlock()

	if lr.started {
		os.Stdout = lr.oldStdout
		_ = lr.w.Close()
		_, _ = lr.buffer.ReadFrom(lr.r)
		lr.started = false
	}
}

func (lr *LogRecorder) ToString() string {
	lr.stop()
	output := lr.buffer.String()
	return RemoveANSISequences(output)
}

func (lr *LogRecorder) ToAnsiString() string {
	lr.stop()
	return lr.buffer.String()
}

func (lr *LogRecorder) ToArray() []string {
	lr.stop()
	output := RemoveANSISequences(lr.buffer.String())
	return strings.FieldsFunc(output, func(r rune) bool {
		return r == '\n' || r == '\r'
	})
}
func (lr *LogRecorder) ToAnsiArray() []string {
	lr.stop()
	output := lr.buffer.String()
	return strings.Split(strings.TrimSuffix(output, "\n"), "\n")
}
