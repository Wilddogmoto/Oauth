package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

type (
	writeHook struct { //отправлять любое количество уровней логирования в любой писатель
		Writer   []io.Writer
		LogLevel []logrus.Level
	}
)

func (hook *writeHook) Fire(entry *logrus.Entry) error {

	line, err := entry.String()
	if err != nil {
		fmt.Printf("func Fire Strings error: %v", err)
		return err
	}

	for _, w := range hook.Writer { //перебераем входные писатели
		_, err = w.Write([]byte(line))
		if err != nil {
			fmt.Printf("Writer error: %v", err)
			return err
		}
	}
	return nil
}

func (hook *writeHook) Levels() []logrus.Level { // возвращает уровни из хука
	return hook.LogLevel
}

var Logger *logrus.Entry

func InitLogger() {

	var (
		logger = logrus.New()
	)

	logger.SetReportCaller(true)

	logger.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) { //внутри происходит логирование
			filename := path.Base(frame.File)                                                    // информация о файле
			return fmt.Sprintf("%s", frame.Function), fmt.Sprintf("%s %d", filename, frame.Line) //название файла, номер строчки
		},
	}

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logger.Errorf("error opening the log file: %v", err)
		return
	}

	logger.SetOutput(io.Discard)

	logger.AddHook(&writeHook{
		Writer:   []io.Writer{file, os.Stdout},
		LogLevel: logrus.AllLevels,
	})

	logger.SetLevel(logrus.TraceLevel)

	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	Logger = logrus.NewEntry(logger)
}

