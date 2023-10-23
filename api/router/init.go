package router

import "go.uber.org/fx"

var Init = fx.Options(
	fx.Provide(
		NewRouter,
	),
)
