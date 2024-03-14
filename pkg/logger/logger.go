package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger() *logrus.Logger {
    logger := logrus.New()
    logger.SetFormatter(new(logrus.JSONFormatter))

    // Запись логов в файл
    file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        logger.Fatalf("error opening file: %s", err)
    }
    logger.SetOutput(file)

    return logger
}