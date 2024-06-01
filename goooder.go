package goooder

import (
	"fmt"
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

func (s SeedExecutor) logger() gdconfig.Logger {
	return s.Config.Logger()
}

func (s SeedExecutor) RunWith(tx *sqlx.Tx, name ...string) {
	_name := ""
	if len(name) > 0 {
		_name = name[0]
	}

	for _, seed := range s.Seeders() {
		seedName := fmt.Sprintf("%T", seed)
		if _name == "" || strings.HasSuffix(seedName, _name) {
			s.logger().Infof("run seed:\t%s\n", seedName)

			err := seed.Exec(tx)
			if err != nil {
				s.logger().Errorf("%s\n", err.Error())
				tx.Rollback()
				panic(err)
			}
		}
	}
	tx.Commit()
}

func (s SeedExecutor) Run(name ...string) {
	db, err := sqlx.Connect("postgres", s.Connstr())
	defer db.Close()
	if err != nil {
		s.logger().Fatalf(err.Error())
	}

	tx := db.MustBegin()
	s.RunWith(tx, name...)
}
