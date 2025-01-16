package models

import (
	"time"

	"github.com/AkmalArifin/movie-reservation/internal/db"
	"github.com/guregu/null/v5"
)

type User struct {
	ID        int64       `json:"id"`
	Name      null.String `json:"name"`
	Phone     null.String `json:"phone"`
	Email     null.String `json:"email"`
	Password  null.String `json:"password"`
	Role      null.String `json:"role"`
	CreatedAt null.Time   `json:"created_at"`
	UpdatedAt null.Time   `json:"updated_at"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users(name, phone, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		`

	u.CreatedAt.SetValid(time.Now())
	u.UpdatedAt.SetValid(time.Now())

	_, err := db.DB.Exec(query, u.Name, u.Phone, u.Email, u.Password, u.Role, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}

	query = `SELECT currval('users_id_seq')`
	err = db.DB.QueryRow(query).Scan(&u.ID)

	return err
}

func (u *User) Update() error {
	query := `
		UPDATE users
		SET name = $2, phone = $3, email = $4, password = $5, role = $6, updated_at = $7
		WHERE id = $1`

	u.UpdatedAt.SetValid(time.Now())

	_, err := db.DB.Exec(query, u.ID, u.Name, u.Phone, u.Email, u.Password, u.Role, u.UpdatedAt)

	return err
}

func (u *User) Delete() error {
	query := `
		DELETE FROM users
		WHERE id = $1`

	_, err := db.DB.Exec(query, u.ID)

	return err
}

func GetAllUsers() ([]User, error) {
	query := `
		SELECT id, name, phone, email, password, role, created_at, updated_at
		FROM users`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(id int64) (User, error) {
	query := `SELECT id, name, phone, email, password, role, created_at, updated_at FROM users WHERE id = $1`

	var user User
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Phone, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	query := `SELECT id, name, phone, email, password, role, created_at, updated_at FROM users WHERE email = $1`

	var user User
	row := db.DB.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Phone, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
