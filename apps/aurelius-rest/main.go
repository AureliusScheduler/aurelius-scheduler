package main

import (
	"aurelius/apps/aurelius-rest/bundler"
	"go.uber.org/fx"
)

func main() {
	fx.New(bundler.Module).Run()
}
