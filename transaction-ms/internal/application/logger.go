package application

type Logger interface {
	Error(err error)
	Info(msg string, args ...any)
}
