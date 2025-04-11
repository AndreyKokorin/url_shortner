package links

import (
	"errors"
	"log/slog"
)

func (ls *LinkService) GetToRedirectURL(shortLink string) (string, error) {
	if shortLink == "" {
		return "", errors.New("No short link provided")
	}

	link, err := ls.rep.GetOriginalUrlByShortLink(shortLink)
	if err != nil {
		return "", err
	}

	err = ls.rep.IncrementClickCount(link.Id)
	if err != nil {
		slog.Error("Error incrementing click count", err)
	}

	return link.OriginalURL, nil
}
