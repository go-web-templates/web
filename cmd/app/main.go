package main

import (
	"context"

	"github.com/go-web-templates/web/cmd/app/controllers"
	"github.com/go-web-templates/web/cmd/app/views"
	"github.com/go-web-templates/web/cmd/app/webserver"
	"github.com/go-web-templates/web/internal/application/mappers"
	appservices "github.com/go-web-templates/web/internal/application/services"
	cacherepositories "github.com/go-web-templates/web/internal/infra/cache-repositories"
	"github.com/go-web-templates/web/internal/infra/data"
	eventhandlers "github.com/go-web-templates/web/internal/infra/event-handlers"
	"github.com/go-web-templates/web/internal/infra/repositories"
	infraservices "github.com/go-web-templates/web/internal/infra/services"
	"github.com/go-web-templates/web/pkg/conf"
	"github.com/go-web-templates/web/pkg/logger"
	"go.uber.org/fx"
)

func RegisterWebServer(lc fx.Lifecycle, ws webserver.WebServer) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go ws.StartServer()
			return nil
		},
		OnStop: func(_ context.Context) error {
			go ws.ShutdownServer()
			return nil
		},
	})
}

func main() {
	app := fx.New(
		conf.Module,
		logger.Module,
		data.Module,
		repositories.Module,
		cacherepositories.Module,
		mappers.Module,
		infraservices.Module,
		eventhandlers.Module,
		appservices.Module,
		views.Module,
		controllers.Module,
		webserver.Module,
		fx.Invoke(RegisterWebServer),
	)
	app.Run()
}
