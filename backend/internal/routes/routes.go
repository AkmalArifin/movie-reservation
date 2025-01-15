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
}
