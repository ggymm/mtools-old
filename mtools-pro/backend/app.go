package backend

import (
	"context"

	"github.com/pkg/errors"

	"mtools-pro/backend/pkg/config"
	"mtools-pro/backend/pkg/database"
	"mtools-pro/backend/pkg/log"
	"mtools-pro/backend/service"
)

type App struct {
	service []any
}

func NewApp() *App {
	config.Init()
	return &App{}
}

func (app *App) Init() error {
	log.Init()
	return database.Init()
}

func (app *App) Services() []any {
	if app.service == nil {
		app.service = []any{
			service.NewConvService(),
			service.NewCryptoService(),
			service.NewQrcodeService(),

			service.NewDocService(),
			service.NewScriptService(),
			service.NewSnippetService(),

			service.NewVideoService(),
			service.NewMagnetService(),

			service.NewRuntimeService(), // 通用运行时服务
		}
	}
	return app.service
}

func (app *App) ServicesStartup(ctx context.Context) {
	for _, s := range app.Services() {
		if fn, ok := s.(interface{ Startup(context.Context) }); ok {
			fn.Startup(ctx)
		}
	}
}

func (app *App) OnStartup(ctx context.Context) {
	err := app.Init()
	if err != nil {
		log.Fatal().
			Err(errors.WithStack(err)).
			Msg("startup init storage error")
	}

	app.ServicesStartup(ctx)
}

func (app *App) OnDomReady(_ context.Context) {
}

func (app *App) OnShutdown(_ context.Context) {
}

func (app *App) OnBeforeClose(_ context.Context) bool {
	return false
}
