package seed

import (
	"github.com/AkmalArifin/movie-reservation/internal/models"
	"github.com/AkmalArifin/movie-reservation/internal/utils"
)

func seederUsers() error {
	var userNames = []string{
		"admin",
		"user",
	}
	var userPhones = []string{
		"0",
		"0",
	}
	var userEmails = []string{
		"admin@example.com",
		"user@example.com",
	}
	var userPasswords = []string{
		"admin",
		"12345678",
	}
	var userRoles = []string{
		"admin",
		"user",
	}

	for i := range userNames {
		var user models.User
		user.Name.SetValid(userNames[i])
		user.Phone.SetValid(userPhones[i])
		user.Email.SetValid(userEmails[i])
		user.Role.SetValid(userRoles[i])

		password, err := utils.GeneratePassword(userPasswords[i])
		if err != nil {
			return err
		}

		user.Password.SetValid(password)

		err = user.Save()
		if err != nil {
			return err
		}
	}

	return nil
}
