package links

import "errors"

func (ls *LinkService) DeleteLink(linkID int, userID int) error {
	link, err := ls.rep.GetLinkByID(linkID)
	if err != nil {
		return err
	}

	if link.UserId != userID {
		return errors.New("User doesn't have the permission to delete the link")
	}

	err = ls.rep.DeleteLink(linkID)
	if err != nil {
		return err
	}

	return nil
}
