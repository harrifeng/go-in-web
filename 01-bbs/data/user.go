package data

import (
	"fmt"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (user *User) Create() (err error) {
	statement := `insert into users (uuid, name, email, password, created_at) values (?,?,?,?,?)`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	id, err := stmt.Exec(createUUID(), user.Name, user.Email, user.Password, time.Now())
	fmt.Println(id)
	return
}
