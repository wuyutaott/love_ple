package data

import (
	"time"
)

type Member struct {
	Id        int
	Uuid      string
	CardId    string
	Phone     string
	Name 	  string
	CreatedAt time.Time
}

func (m *Member) Create() (err error) {
	statement := "insert into members (uuid, card_id, phone, name, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), m.CardId, m.Email, m.Phone, m.Name, time.Now()).Scan(&m.Id, &m.Uuid, &m.CreatedAt)
	return
}

// // Delete user from database
// func (user *User) Delete() (err error) {
// 	statement := "delete from users where id = $1"
// 	stmt, err := Db.Prepare(statement)
// 	if err != nil {
// 		return
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(user.Id)
// 	return
// }

// // Update user information in the database
// func (user *User) Update() (err error) {
// 	statement := "update users set name = $2, email = $3 where id = $1"
// 	stmt, err := Db.Prepare(statement)
// 	if err != nil {
// 		return
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(user.Id, user.Name, user.Email)
// 	return
// }

// // Delete all users from database
// func UserDeleteAll() (err error) {
// 	statement := "delete from users"
// 	_, err = Db.Exec(statement)
// 	return
// }

// // Get all users in the database and returns it
// func Users() (users []User, err error) {
// 	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
// 	if err != nil {
// 		return
// 	}
// 	for rows.Next() {
// 		user := User{}
// 		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
// 			return
// 		}
// 		users = append(users, user)
// 	}
// 	rows.Close()
// 	return
// }

// // Get a single user given the email
// func UserByEmail(email string) (user User, err error) {
// 	user = User{}
// 	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
// 		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
// 	return
// }

// // Get a single user given the UUID
// func UserByUUID(uuid string) (user User, err error) {
// 	user = User{}
// 	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = $1", uuid).
// 		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
// 	return
// }
