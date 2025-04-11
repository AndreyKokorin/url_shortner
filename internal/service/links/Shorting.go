package links

import (
	"URL_shortner/internal/model"
	"URL_shortner/pkg/helps"
	"errors"
	"log/slog"
)

func (ls *LinkService) Shorten(userID int, originalUrl string) (*model.Link, error) {
	if originalUrl == "" {
		return nil, errors.New("Invalid URL")
	}

	var shortURL, err = helps.GenerateRandomString(6)

	link := &model.Link{
		UserId:      userID,
		ShortUrl:    shortURL,
		OriginalURL: helps.AddHTTPS(originalUrl),
	}

	linkID, err := ls.rep.NewLink(link)
	if err != nil {
		return nil, err
	}
	slog.Info("shorten func - " + link.OriginalURL)
	slog.Info("shorten func - " + link.ShortUrl)

	link.Id = linkID

	slog.Info("shorten func - ")

	return link, nil
}
