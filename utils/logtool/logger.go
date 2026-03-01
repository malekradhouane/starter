package logtool

import (
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// NewLogger send a new instance of customized logger
func NewLogger(prefix string) *log.Logger {
	return log.New(os.Stderr, prefix+time.Now().Format(" 2006-01-02 - 15:04:05 | "), 0)
}

func SetupLogger(prefix string) *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	log.SetReportCaller(true) // show the logrus line and file
	log.SetFormatter(&PrefixFormatter{
		Prefix: prefix,
	})
	return log
}

type PrefixFormatter struct {
	Prefix string
}

func (f *PrefixFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 - 15:04:05")
	levelColor := getColorByLevel(entry.Level)
	prefixedMessage := levelColor + f.Prefix + " " + timestamp + " | " + entry.Message + "\n" + resetColor()
	return []byte(prefixedMessage), nil
}

func getColorByLevel(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel:
		return "\x1b[36m" // Cyan color
	case logrus.InfoLevel:
		return "\x1b[32m" // Green color
	case logrus.WarnLevel:
		return "\x1b[33m" // Yellow color
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return "\x1b[31m" // Red color
	default:
		return ""
	}
}

func resetColor() string {
	return "\x1b[0m" // Reset color
}
