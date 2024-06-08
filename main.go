package main

import (
	"enfermeria_go/cmd/api"
	"enfermeria_go/config"
	db2 "enfermeria_go/db"
	"log"
)

func main() {

	db, err := db2.NewPostgreSQLStorage(db2.Config{
		Host:     config.Envs.Host,
		Port:     config.Envs.Port,
		User:     config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		DBName:   config.Envs.DBName,
	})

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
