package repository

import (
	"fmt"
	"strconv"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

//digunakan untuk get data sekaligus jika datanya kosng maka akan di buatkan
func (u *UserRepository) LoadOrCreate() ([]User, error) {
	// return []User{}, nil // TODO: replace this
	records, err := u.db.Load("users")
	if err != nil {
		records = [][]string{
			{"username", "password", "loggedin", "role"},
		}
		if err := u.db.Save("users", records); err != nil {
			return nil, err
		}
	}

	result := make([]User, 0)
	for i := 1; i < len(records); i++ {
		loggedIn, err := strconv.ParseBool(records[i][2])
		if err != nil {
			return nil, err
		}

		user := User{
			Username: records[i][0],
			Password: records[i][1],
			Loggedin: loggedIn,
			Role:     records[i][3],
		}
		result = append(result, user)
	}

	fmt.Println(result)
	return result, nil
}

//mengambil semua data
func (u *UserRepository) SelectAll() ([]User, error) {
	// return []User{}, nil // TODO: replace this
	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	// return nil, nil // TODO: replace this
	// users, err := u.SelectAll()
	// if err != nil {
	// 	return nil, err
	// }
	// for _, user := range users {
	// 	if user.Username == username && user.Password == password {
	// 		return &username, nil
	// 	}
	// }
	// return nil, fmt.Errorf("Login Failed")
	records, _ := u.db.Load("users")
	updated := [][]string{
		{"username", "password", "loggedin", "role"},
	}
	var isUserandPassExist bool
	var result *string
	for i := 1; i < len(records); i++ {
		records[i][2] = "false"
		if records[i][0] == username && records[i][1] == password {
			records[i][2] = "true"
			isUserandPassExist = true
			result = &records[i][0]
		}
		updated = append(updated, []string{
			records[i][0],
			records[i][1],
			records[i][2],
			records[i][3],
		})
	}
	if isUserandPassExist == false {
		return nil, fmt.Errorf("Login Failed")
	}

	err := u.db.Save("users", updated)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserRepository) Save(users []User) error {
	// return nil // TODO: replace this
	records := [][]string{
		{"username", "password", "loggedin", "role"},
	}

	for i := 0; i < len(users); i++ {
		records = append(records, []string{
			users[i].Username,
			users[i].Password,
			strconv.FormatBool(users[i].Loggedin),
			users[i].Role,
		})
	}
	return u.db.Save("users", records)
}

func (u *UserRepository) GetUserRole(username string) (*string, error) {
	// TODO: answer here
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Username == username {
			return &user.Role, nil
		}
	}
	return nil, fmt.Errorf("Failed to get user role")
}
