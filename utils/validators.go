package utils

import (
	"log"

	"github.com/gookit/validate"
)

type Result struct {
	ID int
}

func AddValidators() {
	validate.AddValidators(validate.M{
		"unique": func(val string, table string, field string, ignoreField string, ignoreValue string) bool {

			var count int

			if ignoreField != "nil" {
				err := DB.QueryRow("SELECT COUNT(id) FROM " + table + " WHERE " + field + " = " + val + " AND " + ignoreField + " != " + ignoreValue).Scan(&count)
				if err != nil {
					log.Fatalf("Unique validator failed with ingnore field")
				}

				return count <= 0
			}

			err := DB.QueryRow("SELECT COUNT(id) FROM " + table + " WHERE " + field + " = " + val).Scan(&count)
			if err != nil {
				log.Fatalf("Unique validator failed")
			}

			return count <= 0
		},
		"exists": func(val string, table string, field string) bool {

			var count int

			err := DB.QueryRow("SELECT COUNT(id) FROM " + table + " WHERE " + field + " = " + val).Scan(&count)
			if err != nil {
				log.Fatalf("Unique validator failed")
			}

			return count > 0
		},
	})

	validate.AddGlobalMessages(map[string]string{
		"exists": "{field} doesn't exists",
		"unique": "{field} has to be unique. Already used",
	})
}
