package repository

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	//return []User{}, nil // TODO: replace this
	data, err := u.db.Load("users")
	if err != nil {
		var record = [][]string{{
			"username", "password", "loggedin",
		}}

		if err := u.db.Save("users", record); err != nil {
			return nil, err
		}
	}

	var users []User

	for i, dat := range data {
		if i != 0 {
			user := User{
				Username: dat[0],
				Password: dat[1],
			}

			if dat[2] == "true" {
				user.Loggedin = true
			} else {
				user.Loggedin = false
			}

			users = append(users, user)
		}
	}

	return users, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	//return []User{}, nil // TODO: replace this
	return u.LoadOrCreate()
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	//return nil, nil // TODO: replace this
	if err := u.LogoutAll(); err != nil {
		return nil, err
	}

	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			u.changeStatus(username, true)
			return &username, nil
		}
	}

	return nil, errors.New("Login Failed")
}

func (u *UserRepository) Save(users []User) error {
	//return nil // TODO: replace this
	record := [][]string{{
		"username", "password", "loggedin",
	}}

	for _, user := range users {
		newRow := []string{
			user.Username, user.Password,
		}

		if user.Loggedin {
			newRow = append(newRow, "true")
		} else {
			newRow = append(newRow, "false")
		}

		record = append(record, newRow)
	}

	return u.db.Save("users", record)
}

func (u *UserRepository) GetUserRole(username string) (*string, error) {
	// TODO: answer here
	
}
