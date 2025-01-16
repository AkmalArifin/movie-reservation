package seed

import "github.com/AkmalArifin/movie-reservation/internal/models"

func seederGenres() error {
	var genreTitles = []string{
		"Action",
		"Comedy",
		"Drama",
		"Fantasy",
		"Science Fiction",
	}

	var genreDescriptions = []string{
		"Thrilling stories filled with intense battles, daring stunts, and heroic feats.",
		"Guaranteed laughs with humorous situations, witty dialogue, and funny characters.",
		"Deep and emotional narratives exploring life's challenges and relationships.",
		"Magical adventures featuring enchanted lands, mystical creatures, and epic tales.",
		"Explore futuristic worlds, advanced technologies, and interstellar adventures.",
	}

	for i := range genreTitles {
		var genre models.Genre
		genre.Title.SetValid(genreTitles[i])
		genre.Description.SetValid(genreDescriptions[i])

		err := genre.Save()
		if err != nil {
			return err
		}
	}

	return nil
}
