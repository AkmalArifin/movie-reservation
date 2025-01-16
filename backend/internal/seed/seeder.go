package seed

func Seeder() error {
	err := seederGenres()
	if err != nil {
		return err
	}

	err = seederUsers()
	return err
}
