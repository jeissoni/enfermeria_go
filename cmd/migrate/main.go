package main

import (
	"database/sql"
	_ "database/sql"
	db2 "enfermeria_go/db"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {

	stringConection := db2.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "test",
		Password: "password",
		DBName:   "mydb",
	}

	db, err := sql.Open("postgres", stringConection.ConnString())

	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(2)

	cmd := os.Args[(len(os.Args) - 1)]
	switch cmd {
	case "up":
		log.Println(3)
		err := m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("Error al crear la migracion: ", err)
		}

	case "down":
		err := m.Down()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}

	//m.Up()

}
