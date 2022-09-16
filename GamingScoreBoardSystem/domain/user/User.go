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
		log.Println(err.Error())
		return nil, err
	}
	defer db1.Close()
	queryBuilder := squirrel.Select("user_id", "name", "score").From("user_info").
		OrderBy("score desc").Limit(5)
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	rows, err := db1.QueryContext(ctx, query, qargs...)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	var usrs []User
	for rows.Next() {
		usr := User{}
		if err := rows.Scan(&usr.UserId, &usr.UserName, &usr.Score); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		usrs = append(usrs, usr)
	}

	return usrs, nil

}
func GetUserById(ctx context.Context, id int) (*User, error) {
	db1, err := db.GetMySqlDBConnection()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer db1.Close()
	queryBuilder := squirrel.Select("user_id", "name", "phone", "email", "score").
		From("user_info").Where(squirrel.Eq{"user_id": id})
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	rows := db1.QueryRow(query, qargs...)
	usr := User{}
	if err = rows.Scan(&usr.UserId, &usr.UserName, &usr.phone, &usr.email, &usr.Score); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &usr, nil

}

func RegisterUser(ctx context.Context, usr *requests.RegisterUserRequest) (int, error) {
	db1, err := db.GetMySqlDBConnection()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	defer db1.Close()
	queryBuilder := squirrel.Insert("user_info").Columns("name", "email", "phone").Values(usr.Name, usr.Email, usr.Phone)
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	res, err := db1.ExecContext(ctx, query, qargs...)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	id, _ := res.LastInsertId()
	return int(id), nil

}
func UpdateUserScores(ctx context.Context, records map[int]int) error {
	db1, err := db.GetMySqlDBConnection()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer db1.Close()
	var userIds []int
	caseStmt := squirrel.Case()
	for id, scr := range records {
		userIds = append(userIds, id)
		caseStmt = caseStmt.When(squirrel.Eq{"user_id": id}, squirrel.Expr("score+?", scr))
	}
	queryBuilder := squirrel.Update("user_info").
		Set("score", caseStmt).
		Where(squirrel.Eq{"user_id": userIds})
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println(query, qargs)
	_, err = db1.ExecContext(ctx, query, qargs...)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil

}
