package routes

import (
	"net/http"
	"strconv"

	"github.com/AkmalArifin/movie-reservation/internal/models"
	"github.com/AkmalArifin/movie-reservation/internal/utils"
	"github.com/gin-gonic/gin"
)

func getAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch data"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not bind JSON"})
		return
	}

	retrievedUser, err := models.GetUserByEmail(user.Email.String)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	isValid := utils.VerifyPassword(user.Password.String, retrievedUser.Password.String)
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	jwtToken, err := utils.GenerateToken(retrievedUser.ID, retrievedUser.Role.ValueOrZero())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(retrievedUser.ID, retrievedUser.Role.ValueOrZero())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "authentication complete", "token": jwtToken, "refreshToken": refreshToken})
}

func register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not bind JSON"})
		return
	}

	hashedPassword, err := utils.GeneratePassword(user.Password.String)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not bind JSON"})
		return
	}

	user.Password.SetValid(hashedPassword)
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not save data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// TODO: update password
func updateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse param"})
		return
	}

	retrievedID := c.GetInt64("id")
	if id != retrievedID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	retrievedUser, err := models.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "account not found"})
		return
	}

	var user models.User
	c.ShouldBindJSON(&user)

	user.ID = id
	user.Password.SetValid(retrievedUser.Password.String)
	user.CreatedAt.SetValid(retrievedUser.CreatedAt.Time)
	err = user.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update data"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "data updated", "user": user})
}

func deleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse param"})
		return
	}

	retrievedID := c.GetInt64("id")
	if id != retrievedID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	user, err := models.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "account not found"})
		return
	}

	err = user.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete data"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "data deleted"})
}
