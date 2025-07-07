package main

import (
	"database/sql"
	"log"

	"github.com/Dante983/social-gopher/cmd/api"
	"github.com/Dante983/social-gopher/config"
	"github.com/Dante983/social-gopher/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMYSQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("Database connection established successfully")
}
