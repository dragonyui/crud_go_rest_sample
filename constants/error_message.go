package constants

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
)

const ERROR_NOT_FOUND  = "Product Not Found !"
const ERROR_DUP_INDEX = "Product Already Exist !"


func ErrorMessageSwitcher(err error)(error){
	switch err {
	case sql.ErrNoRows:
		return errors.New(ERROR_NOT_FOUND)
		
	default:
		switch err.(*mysql.MySQLError).Number {
		case 1062:
			return errors.New(ERROR_DUP_INDEX)
		default:
			return err
		}


	}
}
