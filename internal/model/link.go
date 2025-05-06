package model

import "time"

type Link struct {
	Id            int       `json:"id"`
	OriginalURL   string    `json:"original_url"`
	ShortUrl      string    `json:"short_url"`
	Clicks        int       `json:"clicks"`
	UserId        int       `json:"user_id"`
	Last_click_at time.Time `json:"last_time_redirect"`
}

type LinkRepository interface {
	NewLink(link *Link) (int, error)
	GetLinkByID(linkId int) (*Link, error)
	GetUserLinks(userId int) ([]Link, error)
	GetOriginalUrlByShortLink(shortLink string) (*Link, error)
	DeleteLink(linkId int) error
	IncrementClickCount(linkId int) error
	GetLinkByOriginalUrl(originalUrl string) (*Link, error)
	UpdateOriginalLink(newOriginalLink string, linkID int) (*Link, error)
	ChangeOriginalLink(newOriginalLink string, linkID int) (*Link, error)
}

type LinkService interface {
	Shorten(userID int, originalUrl string) (*Link, error)
	DeleteLink(linkID, userID int) error
	GetUserLinks(userId int) ([]Link, error)
	ChangeOriginalUrl(linkId int, originalUrl string) (*Link, error)
	GetToRedirectURL(shortLink string) (string, error)
	ChengeOriginalLink(newOriginalLink string, linkID int) (*Link, error)
}
