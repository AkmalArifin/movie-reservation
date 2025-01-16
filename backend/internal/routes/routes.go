package routes

import (
	"github.com/AkmalArifin/movie-reservation/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	auth := r.Group("/")
	auth.Use(middleware.Authenticate)

	r.POST("/login", login)
	r.POST("/register", register)

	r.GET("/users", getAllUsers)
	auth.PUT("/users/:id", updateUser)
	auth.DELETE("/users/:id", deleteUser)

	r.GET("/movies", getAllMovies)
	r.GET("/genres", getAllGenres)
	r.GET("/movies/:id", getMovieByID)
	r.GET("/genres/:id", getGenreByID)
	r.GET("/movies-genres", getAllMoviesGenres)
	auth.POST("/movies", createMovie)
	auth.POST("/genres", createGenre)
	auth.PUT("/movies/:id", updateMovie)
	auth.PUT("/genres/:id", updateGenre)
	auth.DELETE("/movies/:id", deleteMovie)
	auth.DELETE("/genres/:id", deleteGenre)
}
