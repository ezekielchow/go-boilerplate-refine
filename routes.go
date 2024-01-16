package main

import (
	"go-boilerplate/controllers"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func WebRoutes(r *gin.Engine) {
	v1 := r.Group("v1")

	userController := new(controllers.UserController)
	user := v1.Group("/users")
	user.GET("", userController.List)

	cc := new(controllers.CategoryController)
	category := v1.Group("/categories")
	category.GET("", cc.GetList)
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
