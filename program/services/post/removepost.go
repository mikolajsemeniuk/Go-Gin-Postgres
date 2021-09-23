package post

import (
	"go-gin-postgres/program/database"
)

const (
	REMOVE_POST_QUERY = "DELETE FROM Posts WHERE id = $1;"
)

func RemovePost(id int64) error {
	_, error := GetPost(id)
	if error != nil {
		return error
	}

	result, error := database.Client.Exec(REMOVE_POST_QUERY, id)
	if error != nil {
		return error
	}

	rows, error := result.RowsAffected()
	if error != nil || rows == 0 {
		return error
	}

	return nil
}
