package catrepository

func (i *sCatRepository) Delete(catId *int) error {
	_, err := i.DB.Exec("DELETE FROM cats WHERE id = $1;", catId)

	if err != nil {
		return err
	}

	return nil
}
