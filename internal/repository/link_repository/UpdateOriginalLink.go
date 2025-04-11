package repository

import (
	"URL_shortner/internal/model"
)

func (pr *PostgresLinkRepository) UpdateOriginalLink(newOriginalLink string, linkID int) (*model.Link, error) {
	var link model.Link

	err := pr.Db.QueryRow("UPDATE links SET original_url = $1 WHERE id=$2 RETURNING id, original_url, short_url,click ,user_id",
		newOriginalLink, linkID).Scan(&link.Id,
		&link.OriginalURL,
		link.ShortUrl,
		&link.Clicks,
		&link.UserId)

	if err != nil {
		return nil, err
	}
	return &link, nil
}
