package postgres

import (
	"app/modles"
	"context"
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
	user := modles.NewUsers()
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