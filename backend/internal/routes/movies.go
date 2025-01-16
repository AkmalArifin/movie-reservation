package routes

import (
	"net/http"
	"strconv"

	"github.com/AkmalArifin/movie-reservation/internal/models"
	"github.com/gin-gonic/gin"
)

func getAllMovies(c *gin.Context) {
	movies, err := models.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch data"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func getMovieByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	movie, err := models.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch data"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func getAllGenres(c *gin.Context) {
	genres, err := models.GetAllGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, genres)
}

func getGenreByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	genre, err := models.GetGenreByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch data"})
		return
	}

	c.JSON(http.StatusOK, genre)
}

func getAllMoviesGenres(c *gin.Context) {
	moviesGenres, err := models.GetAllMoviesGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch data"})
		return
	}

	c.JSON(http.StatusOK, moviesGenres)
}

// TODO: how to create genres movies
func createMovie(c *gin.Context) {

}

func createGenre(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	var genre models.Genre
	err := c.ShouldBindJSON(&genre)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	err = genre.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not save data"})
		return
	}

	c.JSON(http.StatusAccepted, genre)
}

// TODO: how to change genres movie
func updateMovie(c *gin.Context) {

}

func updateGenre(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	_, err = models.GetGenreByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "data not found"})
		return
	}

	var genre models.Genre
	err = c.ShouldBindJSON(&genre)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	genre.ID = id
	err = genre.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update data"})
		return
	}

	c.JSON(http.StatusAccepted, genre)
}

func deleteMovie(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	movie, err := models.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "data not found"})
		return
	}

	err = movie.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete data"})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{})
}

func deleteGenre(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	genre, err := models.GetGenreByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "data not found"})
		return
	}

	err = genre.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete data"})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{})
}
