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
