package bundler

import (
	"aurelius/libs/aurelius-mysql/options"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(options.NewDbOptions),
)

func runApplication() {
	// register in crontab

}
