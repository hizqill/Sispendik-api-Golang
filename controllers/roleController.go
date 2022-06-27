package controllers

import (
	"net/http"
	"sispendik-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RoleUserLogin godoc
// @Summary Get all Roles User Login.
// @Description Get a list of Role User Login.
// @Tags Roles
// @Produce json
// @Success 200 {object} []models.Roles
// @Router /roles [get]
func RoleUserLogin(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var roles []models.Roles
	db.Find(&roles)

	c.JSON(http.StatusOK, gin.H{"data": roles})
}
