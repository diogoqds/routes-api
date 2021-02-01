package seeds

import (
	"log"

	"github.com/diogoqds/routes-challenge-api/infra"
)

func CreateSeeds() error {

	adminInsertSql := "INSERT INTO admins (email) VALUES ('admin@email.com') RETURNING id"
	defaultRouteInsertSql := "INSERT INTO routes (name) VALUES ('Outros') RETURNING id"
	var id int

	err := infra.DB.QueryRow(adminInsertSql).Scan(&id)
	if err != nil {
		log.Println("Error saving the admin: " + err.Error())
		return err

	}
	err = infra.DB.QueryRow(defaultRouteInsertSql).Scan(&id)
	if err != nil {
		log.Println("Error saving the default route: " + err.Error())
		return err
	}

	return nil
}
