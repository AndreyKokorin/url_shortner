package repository

import "URL_shortner/internal/model"

func (pr *PostgresLinkRepository) GetOriginalUrlByShortLink(shortLink string) (*model.Link, error) {
	var link model.Link

	err := pr.Db.QueryRow("SELECT id, original_url, short_url, click, user_id, last_click_at FROM links WHERE short_url = $1", shortLink).Scan(&link.Id, &link.OriginalURL, &link.ShortUrl, &link.Clicks, &link.UserId, &link.Last_click_at)
	if err != nil {
		return nil, err
	}

	return &link, nil
}
