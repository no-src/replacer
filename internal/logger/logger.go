package logger

import (
	"fmt"

	"github.com/no-src/log"
	"github.com/no-src/log/level"
)

// InitLogger init default logger
func InitLogger(name string, logDir string, logLevel level.Level) error {
	fLogger, err := log.NewFileLogger(logLevel, logDir, name)
	if err != nil {
		fmt.Printf("[%s] init file logger error => %s\n", name, logDir)
		return err
	}
	cLogger := log.NewConsoleLogger(logLevel)
	log.InitDefaultLogger(log.NewMultiLogger(cLogger, fLogger))
	return nil
}
