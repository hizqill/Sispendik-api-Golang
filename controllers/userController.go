package controllers

import (
	"net/http"
	"sispendik-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserInput struct {
	RoleID   uint   `json:"role_id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetAllUser godoc
// @Summary Get all Get All Users.
// @Description Get a list of Get All Users.
// @Tags Users
// @Produce json
// @Success 200 {object} []models.Users
// @Router /users [get]
func GetALLUsers(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var users []models.Users
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
