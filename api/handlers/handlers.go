package handlers

import (
	"app/modles"
	"app/storage/postgres"
	"context"
	"fmt"
	"log"
)
 type Handlers struct{
	UserRepo  postgres.UserRepo

 }
 func NewHandlers(Userrepo postgres.UserRepo)Handlers{
	return Handlers{ UserRepo: Userrepo}
 }
func (h *Handlers)CreateUser(UserRepo postgres.UserRepo){
	var user = modles.NewUsers()
	fmt.Println("enter user name ")
	fmt.Scanln(&user.User_name)
	fmt.Println("enter gmail ")
	fmt.Scanln(&user.Gmail)
	err:=UserRepo.CreateUser(context.Background(),user)
	if err!=nil{
		log.Println("error on create user",err)
		return 
	}

	fmt.Println("created")

}