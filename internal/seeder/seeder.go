package seeder

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type SeedExecutor struct {
	command string
	args    []string
	seeders []Seeder
}

type Seeder interface {
	Exec(db *sqlx.DB) error
}

func NewSeedExecutor(command string, args []string, seeders []Seeder) *SeedExecutor {
	return &SeedExecutor{
		command: command,
		args:    args,
		seeders: seeders,
	}
}

func (s SeedExecutor) Run() {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	for _, seed := range s.seeders {
		tx := db.MustBegin()
		name := fmt.Sprintf("%T", seed)
		if s.command == "all" || (s.command == "single" && name == s.args[0]) {
			fmt.Printf("====== %T\n", seed)

			err := seed.Exec(db)
			if err != nil {
				fmt.Printf("Error!: %s\n", err.Error())
				tx.Rollback()
				return
			}
			tx.Commit()
		}
	}
}

func mustLoadEnv() {
	var err error
	var envfile string
	env := os.Getenv("GOOOMA_ENV")
	if "" == env {
		err = godotenv.Load()
		envfile = ".env"
	} else {
		envfile = ".env" + env
		err = godotenv.Load(envfile)
	}

	if err != nil {
		log.Fatalf("Error loading %s\n", envfile)
	}
}
