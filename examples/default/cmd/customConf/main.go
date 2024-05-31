package main

import (
	"github.com/version-1/default/db/seeders/development"
	"github.com/version-1/goooder"
	gdconfig "github.com/version-1/goooder/config"
	"github.com/version-1/goooder/seeder"
)

type config struct {
	seeders []gdconfig.Seeder
}

func (c config) Connstr() string {
	return "postgres://postgres:password@127.0.0.1:54321/example?sslmode=disable"
}

func (c config) Seeders() []gdconfig.Seeder {
	return c.seeders
}

func main() {
	r := seeder.NewRenderer()
	conf := config{
		seeders: development.NewSeeders(r),
	}

	executor := goooder.NewSeedExecutor(conf)
	executor.Run()
}
