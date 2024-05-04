package matchrepository

func (i *sMatchRepository) Reject(id *int) error {
	_, err := i.DB.Exec("UPDATE cat_matches SET status = 'rejected' WHERE id = $1;", id)

	if err != nil {
		return err
	}

	return nil
}
