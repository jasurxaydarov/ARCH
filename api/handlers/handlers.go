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


func (h *Handlers)CreateUser(){
	var user  modles.Users
	fmt.Println("enter user name ")
	fmt.Scanln(&user.User_name)
	fmt.Println("enter gmail ")
	fmt.Scanln(&user.Gmail)
	err:=h.UserRepo.CreateUser(context.Background(),user)
	if err!=nil{
		log.Println("error on create user",err)
		return 
	}

	fmt.Println("created")

}

func (h *Handlers)UpdateUser(){

	id:=0
	user_name:=""

	fmt.Println("enter Updating user's name ")
	fmt.Scanln(&user_name)

	fmt.Println("enter Updating user's id ")
	fmt.Scanln(&id)

	err:=h.UserRepo.UpdateUsersName(context.Background(),user_name,0)
	if err != nil{
		log.Println(err)
		return
	}

}


func (h * Handlers)DeleteUser(){
	id:=0

	fmt.Println("enter deleting user's id ")
	fmt.Scanln(&id)

	err:=h.UserRepo.DeleteUser(context.Background(),id)

	if err != nil{
		log.Println(err)
		return
	}

}




func (h *Handlers)GetUsers(){
	
	page:=0
	fmt.Println("enter page number")
	fmt.Scanln(&page)
	rows,err:=h.UserRepo.GetUsers(context.Background(),3,page)
	if err != nil{
		log.Println(err)
		return
	}

	fmt.Println(rows)
}



func (h *Handlers)GetUserByID(){
	id:=0
	fmt.Println("enter user_id")
	fmt.Scanln(&id)
	query,err:=h.UserRepo.GetUserByID(context.Background(),id)

	if err != nil{
		log.Println(err)
		return
	}
	fmt.Println(query)

}
