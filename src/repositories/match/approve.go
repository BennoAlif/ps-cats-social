package matchrepository

import "log"

func (i *sMatchRepository) Approve(id *int) error {
	_, err := i.DB.Exec("UPDATE cat_matches SET status = 'approved' WHERE id = $1;", id)

	if err != nil {
		log.Printf("Error approve match: %s", err)
		return err
	}

	return nil
}
