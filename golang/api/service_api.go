package api

import (
	"errors"

	"github.com/Xenobyte42/techdb/golang/models"
)

func (self *Database) ClearDatabase() (err error) {
	sqlStatement := `TRUNCATE TABLE `
	for _, table := range tables {
		_, err = self.db.Exec(sqlStatement + table)
		if err != nil {
			return
		}
	}
	return
}

func (self *Database) Count(tableName string) (int64, error) {
	var cnt int64
	sqlStatement := `SELECT COUNT(*) FROM `
	rows, err := self.db.Query(sqlStatement + tableName)
	if err != nil {
		return cnt, err
	}
	defer rows.Close()
	if rows.Next() {
		if err = rows.Scan(&cnt); err != nil {
			return cnt, err
		}
	} else {
		return cnt, errors.New("Count must return 1 row!")
	}
	return cnt, nil
}

func (self *Database) GetInfo() (*models.DatabaseStatus, error) {
	data := &models.DatabaseStatus{}
	cnt, err := self.Count("Forums")
	if err != nil {
		return nil, err
	}
	data.Forum = int32(cnt)

	cnt, err = self.Count("Posts")
	if err != nil {
		return nil, err
	}
	data.Post = cnt

	cnt, err = self.Count("Threads")
	if err != nil {
		return nil, err
	}
	data.Thread = int32(cnt)

	cnt, err = self.Count("Users")
	if err != nil {
		return nil, err
	}
	data.User = int32(cnt)
	return data, nil
}
