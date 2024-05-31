package main

import (
	"github.com/version-1/default/db/seeders/development"
	"github.com/version-1/goooder"
	gdconfig "github.com/version-1/goooder/config"
)

func main() {
	seed := development.NewSeed()

	conf := gdconfig.FromEnv(seed.Seeders())

	executor := goooder.NewSeedExecutor(conf)
	executor.Run()
}
