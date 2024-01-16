package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func connectToDatabase(dsn string) {
	if len(dsn) == 0 {
		log.Fatal("Datasource is not set")
	}

	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}

	DB.SetConnMaxLifetime(-1)
	DB.SetMaxIdleConns(3)
	DB.SetMaxOpenConns(3)

	if err = DB.Ping(); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
}

func SetupDatabase(dsn string) {
	connectToDatabase(dsn)
	initBootstrapTables()
	runMigrations()
}

func initBootstrapTables() {
	_, err := DB.Exec(
		`CREATE TABLE IF NOT EXISTS MIGRATIONS(
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)`)

	if err != nil {
		log.Fatalf("Failed to create migration table %v", err.Error())
	}
}

func migrationExists(needle string, haystack []string) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}

func runMigrations() {

	rows, err := DB.Query("SELECT name FROM migrations")

	if err != nil {
		log.Fatalf("Unable to query migrations %v", err.Error())
	}
	defer rows.Close()

	migrationNames := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		migrationNames = append(migrationNames, name)
	}

	files, err := os.ReadDir("./migrations")

	if err != nil {
		log.Fatalf("Unable to load migrations %v", err.Error())
	}

	for i := 0; i < len(files); i++ {
		file := files[i]

		b, err := os.ReadFile("./migrations/" + file.Name())

		if migrationExists(file.Name(), migrationNames) {
			continue
		}

		if err != nil {
			log.Fatalf("Unable to read migration %v", err.Error())
		}

		_, err = DB.Exec(string(b[:]))

		if err != nil {
			log.Fatalf("Failed to create table %v", err.Error())
		}

		if _, err = DB.Exec("INSERT INTO MIGRATIONS (name) VALUES ('" + file.Name() + "')"); err != nil {
			log.Fatalf("Failed to update migrations table after creating migration %v %v", file.Name(), err.Error())
		}

	}
}
