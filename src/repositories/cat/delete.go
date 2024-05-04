package catrepository

import "log"

func (i *sCatRepository) Delete(catId *int) error {
	_, err := i.DB.Exec("DELETE FROM cats WHERE id = $1;", catId)

	if err != nil {
		log.Printf("Error deleting cat: %s", err)
		return err
	}

	return nil
}
