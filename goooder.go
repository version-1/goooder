package goooder

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/version-1/goooder/config"
)

type SeedExecutor struct {
	config.Config
}

func NewSeedExecutor(config config.Config) *SeedExecutor {
	return &SeedExecutor{
		Config: config,
	}
}

func (s SeedExecutor) Run(name ...string) {
	db, err := sqlx.Connect("postgres", s.Connstr())
	if err != nil {
		log.Fatal(err)
	}

	_name := ""
	if len(name) > 0 {
		_name = name[0]
	}

	tx := db.MustBegin()
	for _, seed := range s.Seeders() {
		seedName := fmt.Sprintf("%T", seed)
		if _name == "" || _name == seedName {
			fmt.Printf("====== %s\n", seedName)

			err := seed.Exec(db)
			if err != nil {
				fmt.Printf("Error!: %s\n", err.Error())
				tx.Rollback()
				return
			}
		}
	}
	tx.Commit()
}
