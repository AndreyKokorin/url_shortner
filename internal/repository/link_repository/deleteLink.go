package repository

import "errors"

func (pr *PostgresLinkRepository) DeleteLink(linkId int) error {
	result, err := pr.Db.Exec("DELETE FROM links WHERE id=$1", linkId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("No rows were affected")
	}

	return nil
}
