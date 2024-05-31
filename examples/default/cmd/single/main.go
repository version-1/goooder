package main

import (
	"github.com/version-1/default/db/seeders/development"
	"github.com/version-1/goooder"
	gdconfig "github.com/version-1/goooder/config"
	"github.com/version-1/goooder/seeder"
)

type config struct{}

func (c config) Connstr() string {
	return "postgres://postgres:password@127.0.0.1:54321/example?sslmode=disable"
}

func (c config) Seeders() []gdconfig.Seeder {
	r := seeder.NewRenderer()
	seeders := development.NewSeeders(r)

	return seeders
}

func main() {
	conf := config{}

	executor := goooder.NewSeedExecutor(conf)
	executor.Run("Seed_0000010_CreateUsers")
}
