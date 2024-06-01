package postgres

import (
	"app/modles"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	db *pgx.Conn
}



func NewUsersRepo(db *pgx.Conn)UserRepo{

	return UserRepo{db: db}
}
func (u *UserRepo) GetUserByID(ctx context.Context, userId int) (modles.Users, error) {
	var user modles.Users
	query := "SELECT * FROM users WHERE user_id=$1"
	err := u.db.QueryRow(ctx, query, userId).Scan(&user.User_id, &user.User_name, &user.Gmail)
	if err != nil {
		log.Println("err on GetUserByID ", err)
		return user, err
	}

	return user,nil
}

func (u UserRepo)CreateUser(ctx context.Context,user modles.Users)error{

	query:="INSERT INTO users(user_name,gmail)VALUES($1,$2)"
	_,err:=u.db.Exec(ctx,query,user.User_name,user.Gmail)
		if err!= nil{

			log.Println("error on CreateUser ", err)
			return err
		}
	return nil
}

func (u *UserRepo) GetUsers(ctx context.Context,limit int, page int) ([]modles.Users, error) {
	
	query := "SELECT * FROM users LIMIT $1 OFFSET $2"
	rows,err := u.db.Query(ctx, query,limit,page )

	var users []modles.Users
	var user modles.Users
 
	for rows.Next(){

		rows.Scan(&user.User_id, &user.User_name, &user.Gmail)
		if err != nil {
			log.Println("err on GetUsers ", err)
			return users, err
		}
		users =append(users, user)
	}
	

	return users,nil
}
func (u UserRepo)UpdateUserName(ctx context.Context,name string,user_id int)error{

		query:="UPDATE users SET user_name = $1 WHERE user_id = $2;"

	_, err := u.db.Exec(ctx,query,name,user_id )

		if err!= nil{

			log.Println("error on Update user name ", err)
			return err
		}

	fmt.Println("succesfully updated")
	return nil
}

func (u UserRepo)UpdateUsersName(ctx context.Context,name string,user_id int)error{
	query:="UPDATE users SET user_name = $1 WHERE user_id = $2;"

	_, err := u.db.Exec(ctx,query,name,user_id )

	if err!= nil{

		log.Println("error on Update user name ", err)
		return err
	}

	fmt.Println("succesfully updated")

	return nil
}

func (u UserRepo) DeleteUser(ctx context.Context,user_id int)error{

	query:= "DELETE  FROM users WHERE user_id=$1"

	_,err:=u.db.Exec(ctx,query,user_id)
	if err!= nil{

		log.Println("error on delete user ", err)
		return err
	}
	fmt.Println("succesfully deleted")
	return nil
}