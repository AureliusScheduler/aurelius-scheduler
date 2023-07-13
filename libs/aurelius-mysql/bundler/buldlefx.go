package bundler

import (
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/repository"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(db_context.NewDbContext),
	repository.Module,
)
