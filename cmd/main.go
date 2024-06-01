package main

import (
	"app/api"
	"app/config"
	"app/storage/postgres"
	"context"

	"fmt"
	"log"
)

func main() {
	cfg := config.NEwConfig()

	conn, err := postgres.Conn(cfg)

	if err != nil {
		log.Println("error on conn", conn)
		return
	}
	defer conn.Close(context.Background())

	fmt.Println(conn)
	UserRepo := postgres.NewUsersRepo(conn)

	api.Api(UserRepo)

}
