package log

import (
	"github.com/ggymm/gopkg/logger"
	"github.com/rs/zerolog"

	"mtools-pro/backend/pkg/config"
)

var log zerolog.Logger

func Init() {
	log = logger.Init(config.AppLogPath)
}
