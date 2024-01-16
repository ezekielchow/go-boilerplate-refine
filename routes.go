package main

import (
	"fmt"
	"go-boilerplate/controllers"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.URL.Path, "asd") {
		fmt.Fprintf(w, "gotcha")
		return
	}

	fmt.Fprintf(w, "came in here")
}

func ApiHandler(r *gin.Engine) {
	api := r.Group("api")
	v1 := api.Group("v1")

	userController := new(controllers.UserController)
	user := v1.Group("/users")
	user.GET("/", userController.List)
}

func WebHandler(router *http.ServeMux) {
	router.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("pages/dashboard.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func StaticHandler(router *http.ServeMux) {
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("js/"))))
}
