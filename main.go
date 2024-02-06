package main

import (
	"database/sql"
	"log"

	"github.com/brianw0924/simplebank/api"
	db "github.com/brianw0924/simplebank/db/sqlc"
	"github.com/brianw0924/simplebank/util"
	_ "github.com/lib/pq" // we don't actually call any function of lib/pq directly so the go fomatter will remove it automtically. We have to use blank identifier to keep it.
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: %w", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
