package main

import (
	"aurelius/apps/aurelius-agent-manager/bundler"
	"go.uber.org/fx"
)

func main() {
	fx.New(bundler.Module).Run()
}
