package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.POST("/login", login)
	r.POST("/register", register)
	r.GET("/users", getAllUsers)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)
}
