package config

import (
	"os"
	"path/filepath"

	"github.com/ggymm/gopkg/sys/files"
)

const (
	WindowTitle     = "工作空间"
	WidowWidth      = 1600
	WindowHeight    = 1000
	MinWindowWidth  = 1280
	MinWindowHeight = 800
)

var (
	AppName    string
	BasePath   string
	AppLogPath string
)

func configPath() string {
	path, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	return path
}

func Init() {
	AppName = "mtools-pro"

	BasePath = filepath.Join(configPath(), AppName)
	if !files.Exist(BasePath) {
		_ = os.MkdirAll(BasePath, os.ModePerm)
	}

	AppLogPath = filepath.Join(BasePath, "logs")
	if !files.Exist(AppLogPath) {
		_ = os.MkdirAll(AppLogPath, os.ModePerm)
	}
	AppLogPath = filepath.Join(AppLogPath, "app.log")
}
