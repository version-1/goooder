package seeder

import (
	"bytes"
	"html/template"

	"github.com/jmoiron/sqlx"
	"github.com/version-1/goooder/config"
)

type Renderer struct {
	tmpl *template.Template
}

func NewRenderer(pathGlob ...string) *Renderer {
	path := "db/seeders/development/templates/*.sql"
	if len(pathGlob) > 0 {
		path = pathGlob[0]
	}

	tmpl := template.Must(template.ParseGlob(path))

	return &Renderer{
		tmpl: tmpl,
	}
}

func (r *Renderer) RenderTemplate(filename string) (string, error) {
	buf := new(bytes.Buffer)
	if err := r.tmpl.ExecuteTemplate(buf, filename, struct{}{}); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (r *Renderer) ExecWithFilename(tx *sqlx.Tx, filename string) error {
	query, err := r.RenderTemplate(filename)
	if err != nil {
		return err
	}

	tx.MustExec(query)
	return nil
}

type Seed struct {
	r       *Renderer
	seeders []config.Seeder
}

func New() *Seed {
	return &Seed{
		r: NewRenderer(),
	}
}

func (s *Seed) Seeders() []config.Seeder {
	return s.seeders
}

func (s *Seed) Append(factory func(r *Renderer) config.Seeder) {
	s.seeders = append(s.seeders, factory(s.r))
}

type Receiver interface {
	With(r *Renderer) config.Seeder
}

func (s *Seed) BatchAppend(factory []Receiver) {
	for _, f := range factory {
		s.seeders = append(s.seeders, f.With(s.r))
	}
}
