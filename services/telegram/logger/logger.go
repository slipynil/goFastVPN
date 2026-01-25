package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MyLogger struct {
	Logger *zap.Logger
}

func (l *MyLogger) IsErr(text string, err error) {
	if err != nil {
		l.Logger.Error(text, zap.Error(err))
	}
}

// logger, closeLogFile, err
func NewLogger() (*MyLogger, func() error, error) {
	// создаем директорию логирования
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, nil, fmt.Errorf("mkdir log folder: %w", err)
	}

	// задаем свою локацию логирования времени
	location, _ := time.LoadLocation("Europe/Moscow")
	myTime := time.Now().In(location)

	// задаем названия логгированных файлов и пишем путь для сохранения файла
	timestamp := myTime.Format("2006-01-02T15-04-05.00000")
	logFilePath := filepath.Join("logs", fmt.Sprintf("%s.log", timestamp))

	// создаем и открываем для записи файл
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("open log file: %w", err)
	}

	// задаем конфигурацию логгирования
	fileCfg := zap.NewDevelopmentEncoderConfig()
	termCfg := zap.NewDevelopmentEncoderConfig()

	fileCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15-04-05.00000")

	fileEncoder := zapcore.NewConsoleEncoder(fileCfg)
	termEncoder := zapcore.NewConsoleEncoder(termCfg)

	// ДВА РАЗНЫХ УРОВНЯ ЛОГИРОВАНИЯ
	// решил сделать 1 уровень т.к. пустой файл без ошибок не канон
	fileLvl := zap.NewAtomicLevelAt(zap.InfoLevel)
	termLvl := zap.NewAtomicLevelAt(zap.InfoLevel)

	// ядро, распределение логера куда он будет писать
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), fileLvl),
		zapcore.NewCore(termEncoder, zapcore.AddSync(os.Stdout), termLvl),
	)

	// сам логгер
	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	myLogger := MyLogger{Logger: logger}

	return &myLogger, logFile.Close, nil
}
