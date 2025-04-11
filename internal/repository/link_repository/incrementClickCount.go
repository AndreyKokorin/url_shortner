package repository

import "errors"

func (pr *PostgresLinkRepository) IncrementClickCount(linkID int) error {
	result, err := pr.Db.Exec("UPDATE links SET click = click + 1 WHERE id = $1", linkID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("ErrLinkNotFound")
	}
	return nil
}
