package repository

import (
	"URL_shortner/internal/model"
	"errors"
)

func (pr *PostgresLinkRepository) GetLinkByOriginalUrl(originalUrl string) (*model.Link, error) {
	if originalUrl == "" {
		return nil, errors.New("originalUrl is empty")
	}

	var link model.Link
	err := pr.Db.QueryRow("SELECT  id, original_url, short_url, click, user_id, last_click_at FROM links WHERE original_url = $1", originalUrl).Scan(&link.Id, &link.OriginalURL, &link.ShortUrl, &link.Clicks, &link.UserId, &link.LastClickAt)
	if err != nil {
		return nil, err
	}

	return &link, nil
}
