package main

import (
	"embed"
	"mtools-pro/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"mtools-pro/backend/pkg/config"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := backend.NewApp()

	if err := wails.Run(&options.App{
		Title:     config.WindowTitle,
		Width:     config.WidowWidth,
		Height:    config.WindowHeight,
		MinWidth:  config.MinWindowWidth,
		MinHeight: config.MinWindowHeight,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind:          app.Services(),
		OnStartup:     app.OnStartup,
		OnDomReady:    app.OnDomReady,
		OnShutdown:    app.OnShutdown,
		OnBeforeClose: app.OnBeforeClose,

		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
	}); err != nil {
		println("error: ", err.Error())
	}
}
