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

func (cs CategoryService) List(start string, end string, sortBy string, sortDirection string) []models.Category {
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

	query := "Select * from " + models.CategoriesTable + " LIMIT " + fmt.Sprint(limitInt) + " OFFSET " + string(start)

	if sortBy != "" && sortDirection != "" {
		query = query + " ORDER BY " + sortBy + " " + strings.ToUpper(sortDirection)
	}

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

	return categories
}
