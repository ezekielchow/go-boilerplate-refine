package services

import (
	"context"
	"errors"
	"fmt"
	"go-boilerplate/models"
	"go-boilerplate/utils"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type CategoryService struct{}

func (cs CategoryService) List(queries url.Values) ([]models.Category, int) {

	start := queries.Get("_start")
	end := queries.Get("_end")
	sortBy := queries.Get("_sort")
	sortDirection := queries.Get("_order")
	filters := utils.ExtractFilters(queries, models.CategoriesSearchable)

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
	countQuery := "Select count(*) from " + models.CategoriesTable
	queryActions := ""

	args := pgx.NamedArgs{}

	if len(filters) > 0 {
		queryActions = queryActions + " WHERE "

		for k, v := range filters {
			queryActions = queryActions + k + " LIKE @" + k + " AND "
			args[k] = "%" + v + "%"
		}

		queryActions = queryActions[:len(queryActions)-4]
	}

	var count int
	countQuery = countQuery + queryActions
	if len(filters) > 0 {
		err = db.QueryRow(context.Background(), countQuery, args).Scan(&count)
	} else {
		err = db.QueryRow(context.Background(), countQuery).Scan(&count)
	}

	if err != nil {
		log.Fatalf("Category Listing: counting failed %v", err.Error())
	}

	if sortBy != "" && sortDirection != "" {
		queryActions = queryActions + " ORDER BY " + strings.ToLower(sortBy) + " " + strings.ToUpper(sortDirection)
	}

	queryActions = queryActions + " LIMIT " + fmt.Sprint(limitInt) + " OFFSET " + string(start)

	var rows pgx.Rows
	query = query + queryActions

	if len(filters) > 0 {
		rows, err = db.Query(context.Background(), query, args)
	} else {
		rows, err = db.Query(context.Background(), query)
	}

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

	return categories, count
}

func (cs CategoryService) Create(data models.Category) models.Category {
	db := utils.DB

	var category models.Category
	err := db.QueryRow(context.Background(), "INSERT INTO "+models.CategoriesTable+" (name) VALUES (@name) RETURNING id,name", pgx.NamedArgs{"name": data.Name}).Scan(&category.ID, &category.Name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			log.Fatalf("Category Creation : failed %v", pgErr.Message)
		}
	}

	return category
}

func (cs CategoryService) GetOne(id string) (models.Category, bool) {
	db := utils.DB

	var category models.Category
	err := db.QueryRow(context.Background(), "SELECT * FROM "+models.CategoriesTable+" WHERE id = @id", pgx.NamedArgs{"id": id}).Scan(&category.ID, &category.Name)
	if err != nil {
		return models.Category{}, false
	}

	return category, true
}

func (cs CategoryService) Update(id string, data models.Category) models.Category {
	db := utils.DB

	args := make(pgx.NamedArgs)
	args["id"] = id

	query := "UPDATE " + models.CategoriesTable + " SET "

	values := reflect.ValueOf(data)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		k := strings.ToLower(types.Field(i).Name)
		v := values.Field(i)

		if k == "id" {
			continue
		}

		query = query + " " + k + "=" + "@" + k + ","
		args[k] = fmt.Sprint(v)
	}
	query = query[:len(query)-1]
	query = query + " where id= @id"

	var category models.Category
	err := db.QueryRow(context.Background(), query, args).Scan(&category.ID, &category.Name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			log.Fatalf("Category Creation : failed %v", pgErr.Message)
		}
	}

	return category
}
