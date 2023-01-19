package mysql_utils

import (
	"fmt"
	"github.com/LucianoLupo95/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError("Error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("Invalid data"))
	}
	return errors.NewInternalServerError("Error processing request")
}
