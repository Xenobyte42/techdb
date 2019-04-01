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
