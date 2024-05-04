package matchrepository

import "log"

func (i *sMatchRepository) Reject(id *int) error {
	_, err := i.DB.Exec("UPDATE cat_matches SET status = 'rejected' WHERE id = $1;", id)

	if err != nil {
		log.Printf("Error rejecting match: %s", err)
		return err
	}

	return nil
}
