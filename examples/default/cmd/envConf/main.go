package main

import (
	"github.com/version-1/default/db/seeders/development"
	"github.com/version-1/goooder"
	gdconfig "github.com/version-1/goooder/config"
	"github.com/version-1/goooder/seeder"
)

func main() {
	r := seeder.NewRenderer()
	seeders := development.NewSeeders(r)

	conf := gdconfig.FromEnv()
	conf.SetSeeders(seeders)

	executor := goooder.NewSeedExecutor(conf)
	executor.Run()
}
