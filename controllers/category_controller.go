package controllers

import (
	"go-boilerplate/services"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

// @Description Listing resource for categories
// @Tags category
// @Product json
// @Param _start query string true "Data queried from"
// @Param _end query string true "Data queried to"
// @Param _sort query string false "Sort by field"
// @Param _order query string false "Sort direction"
// @Success 200 {object} models.Success
// @Failure 500 {object} models.Error
// @Router /categories [get]
func (cc CategoryController) GetList(c *gin.Context) {

	cs := new(services.CategoryService)
	c.JSON(200, cs.List(c.Query("_start"), c.Query("_end"), c.Query("_sort"), c.Query("_order")))
}
