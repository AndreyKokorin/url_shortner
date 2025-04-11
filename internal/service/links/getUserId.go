package links

import "URL_shortner/internal/model"

func (ls *LinkService) GetUserLinks(userId int) ([]model.Link, error) {
	links, err := ls.rep.GetUserLinks(userId)

	if err != nil {
		return nil, err
	}

	return links, nil
}
