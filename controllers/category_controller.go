package controllers

import (
	"fmt"
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
// @Success 200 {object} models.ListSuccess
// @Failure 500 {object} models.Error
// @Router /categories [get]
func (cc CategoryController) GetList(c *gin.Context) {

	cs := new(services.CategoryService)
	data, count := cs.List(c.Query("_start"), c.Query("_end"), c.Query("_sort"), c.Query("_order"))

	c.Writer.Header().Add("access-control-expose-headers", "X-Total-Count")
	c.Writer.Header().Add("X-Total-Count", fmt.Sprint(count))
	c.JSON(200, data)
}
