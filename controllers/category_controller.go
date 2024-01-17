package controllers

import (
	"fmt"
	"go-boilerplate/models"
	"go-boilerplate/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

type CategoryController struct{}

// @Description Listing resource for categories
// @Tags category
// @Produce json
// @Param _start query string true "Data queried from"
// @Param _end query string true "Data queried to"
// @Param _sort query string false "Sort by field"
// @Param _order query string false "Sort direction"
// @Success 200 {object} models.ListSuccess
// @Failure 500 {object} models.Error
// @Router /categories [get]
func (cc CategoryController) GetList(c *gin.Context) {

	cs := new(services.CategoryService)
	data, count := cs.List(c.Request.URL.Query())

	c.Writer.Header().Add("access-control-expose-headers", "X-Total-Count")
	c.Writer.Header().Add("X-Total-Count", fmt.Sprint(count))
	c.JSON(http.StatusOK, data)
}

type createValidation struct {
	Name string `json:"name" validate:"required|string|min_len:1"`
}

// @Description Create resource for category
// @Tags category
// @Produce json
// @Param name body string true "Name of category"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.Error
// @Router /categories [post]
func (cc CategoryController) Create(c *gin.Context) {
	var cv createValidation
	err := c.ShouldBindJSON(&cv)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Create category invalid request"})
		return
	}

	v := validate.Struct(cv)

	if !v.Validate() {
		err = v.Errors
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cs := new(services.CategoryService)
	data := cs.Create(models.Category{Name: cv.Name})

	c.JSON(http.StatusOK, data)
}

// @Description Get one category
// @Tags category
// @Produce json
// @Success 200 {object} models.Category
// @Failure 500 {object} models.Error
// @Router /categories/:id [get]
func (cc CategoryController) GetOne(c *gin.Context) {
	id, hasParam := c.Params.Get("id")

	if !hasParam {
		log.Fatalf("Get One Category id Missing")
	}

	cs := new(services.CategoryService)
	category, found := cs.GetOne(id)

	if !found {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, category)
}

// @Description Update one category
// @Tags category
// @Produce json
// @Param name body string false "Name of category"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.Error
// @Router /categories/:id [patch]
func (cc CategoryController) Update(c *gin.Context) {
	id, hasParam := c.Params.Get("id")

	if !hasParam {
		log.Fatalf("Update One Category id Missing")
	}

	var data models.Category
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Update Category unable to bind json"})
		return
	}

	cs := new(services.CategoryService)
	category := cs.Update(id, data)

	c.JSON(http.StatusOK, category)
}

// @Description Delete one category
// @Tags category
// @Produce json
// @Success 200 {object} boolean
// @Failure 500 {object} models.Error
// @Router /categories/:id [delete]
func (cc CategoryController) Delete(c *gin.Context) {
	fmt.Println("oi!!!")

	id, hasParam := c.Params.Get("id")

	if !hasParam {
		log.Fatalf("Delete One Category id Missing")
	}

	cs := new(services.CategoryService)
	if !cs.Delete(id) {
		c.JSON(http.StatusOK, false)
	}

	c.JSON(http.StatusOK, true)
}
