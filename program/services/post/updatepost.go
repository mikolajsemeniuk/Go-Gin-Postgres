package post

import (
	"go-gin-postgres/program/database"
	"go-gin-postgres/program/inputs"
)

const (
	UPDATE_POST_QUERY = "UPDATE Posts SET Title = $1, Description = $2 WHERE Id = $3;"
)

func UpdatePost(id int64, input inputs.PostInput) error {
	_, error := GetPost(id)
	if error != nil {
		return error
	}

	result, error := database.Client.Exec(UPDATE_POST_QUERY, input.Title, input.Description, id)
	if error != nil {
		return error
	}

	rows, error := result.RowsAffected()
	if error != nil || rows == 0 {
		return error
	}

	return nil
}
