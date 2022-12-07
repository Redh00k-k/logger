package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

const (
	ColorBlack = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
)

type LogConfig struct {
	level       int
	isWriteFile bool
	fc          *LogFileConfig
}

type LogFileConfig struct {
	filepath   string
	timeFormat string
}

func New() *LogConfig {
	lfc := &LogFileConfig{
		filepath:   "./log.txt",
		timeFormat: time.RFC3339,
	}
	return &LogConfig{
		level:       INFO,
		isWriteFile: true,
		fc:          lfc,
	}
}

func (l *LogConfig) SetLogLevel(lv int) {
	l.level = lv
}

func (l *LogConfig) SetWriteFile(wf bool) {
	l.isWriteFile = wf
}

func (l *LogConfig) SetFilePath(path string) {
	l.fc.filepath = path
}

func (l *LogConfig) SetTimeFormat(tf string) {
	l.fc.timeFormat = tf
}

func (l *LogConfig) Trace(input string) {
	l.logging(TRACE, "TRACE", ColorCyan, input)
}

func (l *LogConfig) Debug(input string) {
	l.logging(DEBUG, "DEBUG", ColorBlue, input)
}

func (l *LogConfig) Info(input string) {
	l.logging(INFO, "INFO", ColorGreen, input)
}

func (l *LogConfig) Warn(input string) {
	l.logging(WARN, "WARN", ColorYellow, input)
}

func (l *LogConfig) Error(input string) {
	l.logging(ERROR, "ERROR", ColorRed, input)
}

func (l *LogConfig) Fatal(input string) {
	l.logging(FATAL, "FATAL", ColorMagenta, input)
}

func (l *LogConfig) writeFile(line string) {
	file, err := os.OpenFile(l.fc.filepath, os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return
	}

	file.WriteString(line + "\n")

	file.Close()
}

func (l *LogConfig) logging(lv int, lvt string, colorCode int, input string) {
	if l.level > lv {
		return
	}

	now := time.Now()
	line := fmt.Sprintf("%s \x1b[%dm[%s]\t%s\x1b[0m", now.Format(l.fc.timeFormat), colorCode, lvt, input)
	fmt.Println(line)

	if l.isWriteFile {
		l.writeFile(line)
	}
}
