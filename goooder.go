package goooder

import (
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	gdconfig "github.com/version-1/goooder/config"
)

type SeedExecutor struct {
	gdconfig.Config
}

func NewSeedExecutor(config gdconfig.Config) *SeedExecutor {
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
		if _name == "" || strings.HasSuffix(seedName, _name) {
			fmt.Printf("====== %s\n", seedName)

			err := seed.Exec(tx)
			if err != nil {
				fmt.Printf("Error!: %s\n", err.Error())
				tx.Rollback()
				fmt.Println("@@@@@@@@@@@@@")
				return
			}
		}
	}
	tx.Commit()
}
