package user

import (
	"context"
	"intuitMc/db"
	"intuitMc/requests"
	"log"

	"github.com/Masterminds/squirrel"
)

type User struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	phone    string
	email    string
	Score    int `json:"score"`
}

func GetTopKUserScores(ctx context.Context) ([]User, error) {
	db1, err := db.GetMySqlDBConnection()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db1.Close()
	queryBuilder := squirrel.Select("user_id", "name", "score").From("user_info").
		OrderBy("score desc").Limit(5)
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err := db1.QueryContext(ctx, query, qargs...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var usrs []User
	for rows.Next() {
		usr := User{}
		if err := rows.Scan(&usr.UserId, &usr.UserName, &usr.Score); err != nil {
			log.Println(err)
			return nil, err
		}
		usrs = append(usrs, usr)
	}

	return usrs, nil

}
func GetUserById(ctx context.Context, id int) (*User, error) {
	db1, err := db.GetMySqlDBConnection()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db1.Close()
	queryBuilder := squirrel.Select("user_id", "name", "phone", "email", "score").
		From("user_info").Where(squirrel.Eq{"user_id": id})
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows := db1.QueryRow(query, qargs...)
	usr := User{}
	if err = rows.Scan(&usr.UserId, &usr.UserName, &usr.phone, &usr.email, &usr.Score); err != nil {
		log.Println(err)
		return nil, err
	}
	return &usr, nil

}

func RegisterUser(ctx context.Context, usr *requests.RegisterUserRequest) error {
	db1, err := db.GetMySqlDBConnection()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db1.Close()
	queryBuilder := squirrel.Insert("user_info").Columns("name", "email", "phone").Values(usr.Name, usr.Email, usr.Phone)
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db1.ExecContext(ctx, query, qargs...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
func UpdateUserScores(ctx context.Context, id, score int) error {
	db1, err := db.GetMySqlDBConnection()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db1.Close()
	queryBuilder := squirrel.Update("user_info").Set("score", score).Where(squirrel.Eq{"user_id": id})
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err2 := db1.ExecContext(ctx, query, qargs...)
	if err2 != nil {
		log.Println(err)
		return err
	}
	return nil

}
