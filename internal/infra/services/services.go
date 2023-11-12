package services

import (
	"github.com/go-web-templates/web/internal/application/interfaces"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewPlaygroundJsonValidator,
		fx.As(new(interfaces.JsonValidator)),
	),
)
