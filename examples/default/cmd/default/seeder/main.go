package main

import (
	"flag"

	"github.com/version-1/default/db/seeders/development"
	goooder "github.com/version-1/goooder/internal/seeder"
)

func main() {
	seed := development.NewSeed()
	flag.Parse()
	args := flag.Args()
	var options []string
	if len(args) >= 2 {
		options = args[1:]
	}
	command := args[0]

	executor := goooder.NewSeedExecutor(command, options, seed.Seeders)
	executor.Exec()
}
