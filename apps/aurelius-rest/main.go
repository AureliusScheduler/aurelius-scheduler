package main

import (
	"aurelius/apps/aurelius-rest/bundler"
	_ "aurelius/apps/aurelius-rest/docs"
	"go.uber.org/fx"
)

// @title Aurelius REST API
// @host      localhost:8080
func main() {
	fx.New(bundler.Module).Run()
}
