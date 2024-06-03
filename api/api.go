package api

import (
	"app/api/handlers"
	"app/storage/postgres"
	"fmt"
)
	

func Api(UserRepo postgres.UserRepo){

	h:=handlers.NewHandlers(UserRepo)
	oprt:=-1

	
	for oprt!=0{
		fmt.Println(`
		Enter oprt

		1.Get user By user_id
		2.Get users
		3.Create users
		4.Update user name
		5.Delete user
		0.Exit
				
		`)
		fmt.Scanln(&oprt)

		switch oprt{
		case 1: h.GetUserByID()	
		case 2: h.GetUsers()	
		case 3: h.CreateUser()
		case 4: h.UpdateUser()
		case 5: h.DeleteUser()
		case 0:return
		}
	}



}


