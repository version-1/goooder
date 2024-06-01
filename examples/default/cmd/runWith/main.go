package main

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
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

func (c config) Logger() gdconfig.Logger {
	return &gdconfig.DefaultLogger{}
}

func main() {
	conf := config{}
	ctx := context.Background()

	db, err := sqlx.Connect("postgres", conf.Connstr())
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	executor := goooder.NewSeedExecutor(conf)
	executor.RunWith(tx)
}
