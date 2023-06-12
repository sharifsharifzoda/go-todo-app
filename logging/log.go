package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

type fileHook struct {
	LevelsArr []logrus.Level
	Files     map[logrus.Level]*os.File
}

func (hook *fileHook) Fire(entry *logrus.Entry) error {
	for _, level := range hook.LevelsArr {
		if entry.Level <= level {
			entry.Logger.Out = hook.Files[level]
			break
		}
	}
	return nil
}

func (hook *fileHook) Levels() []logrus.Level {
	return hook.LevelsArr
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func InitLog() {
	logger := logrus.New()
	logger.SetReportCaller(true)

	debugFile, err := os.OpenFile("./logs/debug.logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatal(err)
	}
	//defer debugFile.Close()

	infoFile, err := os.OpenFile("./logs/info.logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatal(err)
	}
	//defer infoFile.Close()

	errorFile, err := os.OpenFile("./logs/error.logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatal(err)
	}
	//defer errorFile.Close()

	logger.AddHook(&fileHook{
		LevelsArr: []logrus.Level{
			logrus.DebugLevel,
			logrus.InfoLevel,
			logrus.ErrorLevel,
		},
		Files: map[logrus.Level]*os.File{
			logrus.DebugLevel: debugFile,
			logrus.InfoLevel:  infoFile,
			logrus.ErrorLevel: errorFile,
		},
	})

	//logger.SetLevel(logrus.DebugLevel)

	e = logrus.NewEntry(logger)
}
