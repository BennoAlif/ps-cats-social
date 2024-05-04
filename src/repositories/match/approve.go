package matchrepository

func (i *sMatchRepository) Approve(id *int) error {
	_, err := i.DB.Exec("UPDATE cat_matches SET status = 'approved' WHERE id = $1;", id)

	if err != nil {
		return err
	}

	return nil
}
