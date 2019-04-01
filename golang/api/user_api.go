package api

import (
	"database/sql"

	"github.com/Xenobyte42/techdb/golang/models"
)

func (self *Database) UserExists(nickname string) (bool, error) {
	var exists bool
	sqlStatement := `SELECT exists (SELECT * FROM Users WHERE nickname=$1)`

	err := self.db.QueryRow(sqlStatement, nickname).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}

func (self *Database) CreateUser(userData *models.UserUpdate, nickname string) error {
	sqlStatement := `INSERT INTO Users (about, email, fullname, nickname)
	VALUES ($1, $2, $3, $4)`

	_, err := self.db.Exec(sqlStatement, userData.About,
		userData.Email, userData.Fullname, nickname)
	return err
}

func (self *Database) GetUserData(data *models.UserModel, nickname string) (bool, error) {
	sqlStatement := `SELECT about, email, fullname, nickname FROM Users WHERE nickname=$1`

	err := self.db.QueryRow(sqlStatement, nickname).Scan(&data.About,
		&data.Email, &data.Fullname, &data.Nickname)
	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		return false, err
	default:
		return true, nil
	}
}

func (self *Database) CheckEmail(data *models.UserUpdate, nickname string) (bool, error) {
	sqlStatement := `SELECT nickname FROM Users WHERE email=$1`
	var selectedNick string

	err := self.db.QueryRow(sqlStatement, data.Email).Scan(&selectedNick)
	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		return false, err
	default:
		return selectedNick != nickname, nil
	}
}

func (self *Database) UpdateData(data *models.UserUpdate, 
	  nickname string, response *models.UserModel) (bool, error) {
	sqlStatement := `UPDATE Users SET email=$1, fullname=$2, about=$3 WHERE nickname=$4
	RETURNING about, email, fullname, nickname
	`

	err := self.db.QueryRow(sqlStatement, data.Email,
		data.Fullname, data.About, nickname).Scan(&response.About,
		&response.Email, &response.Fullname, &response.Nickname)
	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		return false, err
	default:
		return true, nil
	}
}
