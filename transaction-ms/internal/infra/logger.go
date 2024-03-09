package infra

import "github.com/sirupsen/logrus"

type LoggerStdout struct {
	logger *logrus.Logger
}

func NewLoggerStdout(logger *logrus.Logger) LoggerStdout {
	return LoggerStdout{
		logger: logger,
	}
}

func (l LoggerStdout) Error(err error) {
	l.logger.Error(err)
}

func (l LoggerStdout) Info(msg string, args ...any) {
	l.logger.Info(msg, args)
}
