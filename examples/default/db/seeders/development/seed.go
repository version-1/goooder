package development

import (
	"bytes"
	"html/template"
	"math/rand"
	"strconv"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/version-1/goooder"
)

type Seed struct {
	tmpl    *template.Template
	Seeders []goooder.Seeder
}

func NewSeed() *Seed {
	tmpl := template.Must(template.ParseGlob("db/seeders/development/templates/*.sql"))

	return &Seed{
		tmpl: tmpl,
		Seeders: []goooder.Seeder{
			Seed_0000010_CreatePlans{},
			Seed_0000020_CreateStripeItems{},
			Seed_0000030_CreateMaterials{},
		},
	}
}

func (s Seed) renderTemplate(filename string) (string, error) {
	buf := new(bytes.Buffer)
	if err := s.tmpl.ExecuteTemplate(buf, filename, struct{}{}); err != nil {
		return "", err
	}

	return buf.String(), nil
}

type Seed_0000010_CreatePlans struct {
	Seed
}

type Seed_0000020_CreateStripeItems struct {
	Seed
}

type Seed_0000030_CreateMaterials struct {
	Seed
}

func (s Seed_0000010_CreatePlans) Exec(tx *sqlx.DB) error {
	query, err := s.renderTemplate("create_plans.sql")
	if err != nil {
		return err
	}

	tx.MustExec(query)
	return nil
}

func (s Seed_0000020_CreateStripeItems) Exec(tx *sqlx.DB) error {
	obj := map[string][]string{
		"PL-000001": {"prod_Jet9M0T0XgHjxJ", "price_1J1ZN8IoZNZEhIpWRoO6451P"},
		"PL-000002": {"prod_JewIAAP9UBDyPB", "price_1J1cQCIoZNZEhIpWNPWFeaZX"},
		"PL-000003": {"prod_JetfPiXRltPBNz", "price_1J1ZsBIoZNZEhIpWNv2OAPgm"},
		"PL-000004": {"prod_JetdJfJwgH1VrZ", "price_1J1ZqLIoZNZEhIpWBtfbWiUk"},
	}

	query, err := s.renderTemplate("create_stripe_price.sql")
	if err != nil {
		return err
	}

	rows, err := tx.Query("SELECT id, display_id FROM plans")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var displayID string
		rows.Scan(&id, &displayID)

		productID := obj[displayID][0]
		priceID := obj[displayID][1]

		tx.MustExec(query, productID, priceID, id, "Price")
	}

	return nil
}

func (s Seed_0000030_CreateMaterials) Exec(tx *sqlx.DB) error {
	const MATERIAL_TYPE_CHALLENGE = 1000
	const MATERIAL_TYPE_TEXT = 2000
	const MATERIAL_TYPE_QUIZ = 3000

	materials := [][]interface{}{
		{"TXT-" + randNStrings(8) + "-01", MATERIAL_TYPE_TEXT, "javascript-debug", "JavaScriptでのデバッグ技術入門"},
		{"CHL-" + randNStrings(8) + "-01", MATERIAL_TYPE_CHALLENGE, "html-sass-coding", "HTML+CSS(SASS) コーディング (デザインサイト模写)"},
		{"TXT-" + randNStrings(8) + "-01", MATERIAL_TYPE_TEXT, "oop-entry", "オブジェクト指向入門"},
		{"QIZ-" + randNStrings(8) + "-01", MATERIAL_TYPE_QUIZ, "javascript-quiz", "JavaScriptの基本文法クイズ"},
	}

	for _, m := range materials {
		query, err := s.renderTemplate("create_materials.sql")
		if err != nil {
			return err
		}
		var id uuid.UUID
		err = tx.QueryRow(query, m...).Scan(&id)
		if err != nil {
			return err
		}

		query, err = s.renderTemplate("create_material_nodes.sql")
		if err != nil {
			return err
		}

		const MATERIAL = 1000
		tx.MustExec(query, id, "/", MATERIAL)
	}

	return nil
}

func randNStrings(length int) string {
	var str string
	for i := 0; i < length; i++ {
		str = str + strconv.Itoa(rand.Intn(10))
	}
	return str
}
