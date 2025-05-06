package repository

import "URL_shortner/internal/model"

func (pr *PostgresLinkRepository) GetUserLinks(userid int) ([]model.Link, error) {
	rows, err := pr.Db.Query("SELECT id, original_url, short_url, click, user_id,last_click_at  FROM links WHERE user_id = $1", userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []model.Link
	for rows.Next() {
		var link model.Link

		err = rows.Scan(&link.Id, &link.OriginalURL, &link.ShortUrl, &link.Clicks, &link.UserId, &link.Last_click_at)

		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}
