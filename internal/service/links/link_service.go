package links

import "URL_shortner/internal/model"

type LinkService struct {
	rep model.LinkRepository
}

func NewLinkService(rep model.LinkRepository) *LinkService {
	return &LinkService{rep: rep}
}
