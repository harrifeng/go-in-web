package data

import (
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

	_, err = stmt.Exec(createUUID(), user.Name, user.Email, user.Password, time.Now())
	return
}

func (user *User) CreateSession() (session Session, err error) {
	statement := `insert into sessions (uuid, email, user_id, created_at) values (?,?,?,?)`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	id, err := stmt.Exec(createUUID(), user.Email, user.Id, time.Now())

	v, err := id.LastInsertId()
	return SessionById(int(v))
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow(
		`SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?`,
		email,
	).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func SessionById(id int) (session Session, err error) {
	session = Session{}
	err = Db.QueryRow(
		`select id, uuid, email, user_id, created_at from sessions where id = ?`,
		id,
	).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)

	return
}
