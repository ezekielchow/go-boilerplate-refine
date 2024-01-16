package services

import (
	"fmt"
	"go-boilerplate/models"
	"go-boilerplate/utils"
	"log"
	"strconv"
	"strings"
)

type CategoryService struct{}

func (cs CategoryService) List(start string, end string, sortBy string, sortDirection string) ([]models.Category, int) {
	db := utils.DB

	startInt, err := strconv.Atoi(start)
	if err != nil {
		log.Fatalf("Category Listing: Unable to convert _start %v", err.Error())
	}

	endInt, err := strconv.Atoi(end)
	if err != nil {
		log.Fatalf("Category Listing: Unable to convert _end %v", err.Error())
	}

	limitInt := endInt - startInt

	categories := []models.Category{}

	query := "Select * from " + models.CategoriesTable

	if sortBy != "" && sortDirection != "" {
		query = query + " ORDER BY " + strings.ToLower(sortBy) + " " + strings.ToUpper(sortDirection)
	}

	query = query + " LIMIT " + fmt.Sprint(limitInt) + " OFFSET " + string(start)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Category Listing: Unable to query %v", err.Error())
	}

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Fatalf("Failed to scan row %v", err.Error())
		}

		categories = append(categories, category)
	}

	var count int
	err = utils.DB.QueryRow("SELECT COUNT(id) FROM " + models.CategoriesTable).Scan(&count)
	if err != nil {
		log.Fatalf("Category Listing: counting failed %v", err.Error())
	}

	return categories, count
}
