package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// @Summary get a list of users
// @Schemes
// @Description get a list of users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /users [get]
func (uc UserController) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": nil})
}
