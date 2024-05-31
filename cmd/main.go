package main

import (
	"app/config"
	"app/storage/postgres"
	"context"
	"fmt"
	"log"
)

func main() {
	cfg:=config.NEwConfig()
	
	conn,err:=postgres.Conn(cfg)

	if err!=nil{
		log.Println("error on conn",conn)
		return
	}
	fmt.Println(conn)
	users:=postgres.NewUsersRepo(conn)

	query,err:=users.GetUserByID(context.Background(),1)
	if err != nil{
		log.Println(err)
		return
	}
	fmt.Println(query)

}
