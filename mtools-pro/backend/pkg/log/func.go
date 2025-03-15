package log

import (
	"github.com/rs/zerolog"
)

func Trace() *zerolog.Event {
	return log.Trace()
}

func Debug() *zerolog.Event {
	return log.Debug()
}

func Info() *zerolog.Event {
	return log.Info()
}

func Warn() *zerolog.Event {
	return log.Warn()
}

func Error() *zerolog.Event {
	return log.Error().Stack()
}

func Fatal() *zerolog.Event {
	return log.Fatal().Stack()
}

func Panic() *zerolog.Event {
	return log.Panic().Stack()
}
