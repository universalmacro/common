package ulog

type Logger interface {
	Error()
	Warm()
	Info()
	Debug()
}
