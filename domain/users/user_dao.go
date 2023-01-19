package users

import (
	"fmt"
	"github.com/LucianoLupo95/bookstore_users-api/datasources/mysql/users_db"
	"github.com/LucianoLupo95/bookstore_users-api/utils/date_utils"
	"github.com/LucianoLupo95/bookstore_users-api/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists: %s", err.Error()"")
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", user.Email))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}
