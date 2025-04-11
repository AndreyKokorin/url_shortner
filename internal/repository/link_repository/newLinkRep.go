package repository

import (
	"URL_shortner/internal/model"
	"log/slog"
)

func (pr *PostgresLinkRepository) NewLink(link *model.Link) (int, error) {
	var linkId int
	err := pr.Db.QueryRow("INSERT INTO links(original_url, short_url, user_id) VALUES ($1, $2, $3) RETURNING id", link.OriginalURL, link.ShortUrl, link.UserId).Scan(&linkId)

	if err != nil {
		return 0, err
	}

	slog.Info("NewLink func - " + link.OriginalURL)
	slog.Info("NewLink func - " + link.ShortUrl)

	return linkId, nil
}
