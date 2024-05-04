package matchrepository

func (i *sMatchRepository) Delete(id *int) error {
	_, err := i.DB.Exec("DELETE FROM cat_matches WHERE id = $1;", id)

	if err != nil {
		return err
	}

	return nil
}
