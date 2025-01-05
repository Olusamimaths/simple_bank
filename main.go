package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Olusamimaths/simple_bank/api"
	db "github.com/Olusamimaths/simple_bank/db/sqlc"
	"github.com/Olusamimaths/simple_bank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("couldn't not load config ", err)
	}
	log.Printf("DB_DRIVER: %s, DB_SOURCE: %s", os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("couldn't start server", err)
	}
}
