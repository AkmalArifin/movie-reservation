package models

import (
	"time"

	"github.com/AkmalArifin/movie-reservation/internal/db"
	"github.com/guregu/null/v5"
)

type Movie struct {
	ID          int64       `json:"id"`
	Title       null.String `json:"title"`
	Description null.String `json:"description"`
	PosterImage null.String `json:"poster_image"`
	Price       null.Int64  `json:"price"`
	CreatedAt   null.Time   `json:"created_at"`
	UpdatedAt   null.Time   `json:"updated_at"`
}

type Genre struct {
	ID          int64       `json:"id"`
	Title       null.String `json:"title"`
	Description null.String `json:"description"`
	CreatedAt   null.Time   `json:"created_at"`
	UpdatedAt   null.Time   `json:"updated_at"`
}

type MovieGenre struct {
	MovieID   int64     `json:"movie_id"`
	GenreID   int64     `json:"genre_id"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

type MovieWithGenres struct {
	ID          int64       `json:"id"`
	Title       null.String `json:"title"`
	Description null.String `json:"description"`
	PosterImage null.String `json:"poster_image"`
	Price       null.Int64  `json:"price"`
	Genres      []string    `json:"genres"`
	CreatedAt   null.Time   `json:"created_at"`
	UpdatedAt   null.Time   `json:"updated_at"`
}

func (m *Movie) Save() error {
	query := `
		INSERT INTO movies(title, description, poster_image, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	m.CreatedAt.SetValid(time.Now())
	m.UpdatedAt.SetValid(time.Now())

	_, err := db.DB.Exec(query, m.Title, m.Description, m.PosterImage, m.Price, m.CreatedAt, m.UpdatedAt)
	if err != nil {
		return err
	}

	query = `SELECT currval('movies_id_seq')`
	err = db.DB.QueryRow(query).Scan(&m.ID)

	return err
}

func (g *Genre) Save() error {
	query := `
		INSERT INTO genres(title, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4)`

	g.CreatedAt.SetValid(time.Now())
	g.UpdatedAt.SetValid(time.Now())

	_, err := db.DB.Exec(query, g.Title, g.Description, g.CreatedAt, g.UpdatedAt)
	if err != nil {
		return err
	}

	query = `SELECT currval('genres_id_seq')`
	err = db.DB.QueryRow(query).Scan(&g.ID)

	return err
}

func (mg *MovieGenre) Save() error {
	query := `
		INSERT INTO movies_genres(movie_id, genre_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4)`

	mg.CreatedAt.SetValid(time.Now())
	mg.UpdatedAt.SetValid(time.Now())

	_, err := db.DB.Exec(query, mg.MovieID, mg.GenreID, mg.CreatedAt, mg.UpdatedAt)

	return err
}

func (m *Movie) Update() error {
	query := `
		UPDATE movies
		SET title = $2, description = $3, poster_image = $4, price = $5, updated_at = $6
		WWHERE id = $1`

	m.UpdatedAt.SetValid(time.Now())

	_, err := db.DB.Exec(query, m.ID, m.Title, m.Description, m.PosterImage, m.Price, m.UpdatedAt)

	return err
}

func (g *Genre) Update() error {
	query := `
		UPDATE genres
		SET title = $2, description = $3, updated_at = $4
		WWHERE id = $1`

	g.UpdatedAt.SetValid(time.Now())

	_, err := db.DB.Exec(query, g.ID, g.Title, g.Description, g.UpdatedAt)

	return err
}

func (m *Movie) Delete() error {
	query := `
		DEKETE FROM movies
		WHERE id = $1`

	_, err := db.DB.Exec(query, m.ID)

	return err
}

func (g *Genre) Delete() error {
	query := `
		DEKETE FROM genres
		WHERE id = $1`

	_, err := db.DB.Exec(query, g.ID)

	return err
}

func (mg *MovieGenre) Delete() error {
	query := `
		DEKETE FROM movies_genres
		WHERE movie_id = $1 AND genre_id = $2`

	_, err := db.DB.Exec(query, mg.MovieID, mg.GenreID)

	return err
}

func GetAllMovies() ([]Movie, error) {
	query := `SELECT id, title, description, poster_image, price, created_at, updated_at FROM movies`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.PosterImage, &movie.Price, &movie.CreatedAt, &movie.UpdatedAt)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func GetMovieByID(id int64) (Movie, error) {
	query := `SELECT id, title, description, poster_image, price, created_at, updated_at FROM movies WHERE id = $1`

	var movie Movie
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.PosterImage, &movie.Price, &movie.CreatedAt, &movie.UpdatedAt)
	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}

func GetAllGenres() ([]Genre, error) {
	query := `SELECT id, title, description, created_at, updated_at FROM genres`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []Genre
	for rows.Next() {
		var genre Genre
		err = rows.Scan(&genre.ID, &genre.Title, &genre.Description, &genre.CreatedAt, &genre.UpdatedAt)
		if err != nil {
			return nil, err
		}

		genres = append(genres, genre)
	}

	return genres, nil
}

func GetGenreByID(id int64) (Genre, error) {
	query := `SELECT id, title, description, created_at, updated_at FROM movies WHERE id = $1`

	var genre Genre
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&genre.ID, &genre.Title, &genre.Description, &genre.CreatedAt, &genre.UpdatedAt)
	if err != nil {
		return Genre{}, err
	}

	return genre, nil
}

func GetAllMoviesGenres() ([]MovieGenre, error) {
	query := `SELECT movie_id, genre_id, created_at, updated_at FROM movies_genres`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moviesGenres []MovieGenre
	for rows.Next() {
		var movieGenre MovieGenre
		err = rows.Scan(&movieGenre.MovieID, &movieGenre.GenreID, &movieGenre.CreatedAt, &movieGenre.UpdatedAt)
		if err != nil {
			return nil, err
		}

		moviesGenres = append(moviesGenres, movieGenre)
	}

	return moviesGenres, nil
}
