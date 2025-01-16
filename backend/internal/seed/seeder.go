package seed

func Seeder() error {
	var err error
	// err = seederGenres()
	// if err != nil {
	// 	return err
	// }

	err = seederUsers()
	return err
}
