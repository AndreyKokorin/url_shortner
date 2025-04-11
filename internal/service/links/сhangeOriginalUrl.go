package links

import (
	"URL_shortner/internal/model"
	"errors"
)

func (ls *LinkService) ChangeOriginalUrl(linkId int, newOriginalUrl string) (*model.Link, error) {
	if newOriginalUrl == "" {
		return nil, errors.New("newOriginalUrl is empty")
	}

	link, err := ls.rep.UpdateOriginalLink(newOriginalUrl, linkId)
	if err != nil {
		return nil, err
	}

	return link, nil
}
