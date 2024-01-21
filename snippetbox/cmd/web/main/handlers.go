package main

import (
	"bolatdauren.net/snippetbox/pkg/models"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func homePageFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/layout_home.tmpl",
		"./ui/html/layout_base.tmpl",
		"./ui/html/powered_by.tmpl",
	}

	latestArticles, err := models.GetLatestArticles()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	data := struct {
		LatestArticles []models.Article
	}{
		LatestArticles: latestArticles,
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
	}
}

func forStudentsFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/students_path" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/layout_students.tmpl",
		"./ui/html/layout_base.tmpl",
		"./ui/html/powered_by.tmpl",
	}

	articlesForStudents, err := models.GetArticlesByCategory("студентам")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	data := struct {
		LatestArticles []models.Article
	}{
		LatestArticles: articlesForStudents,
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
	}
}

func forStaffFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/staff_path" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/layout_staff.tmpl",
		"./ui/html/layout_base.tmpl",
		"./ui/html/powered_by.tmpl",
	}

	articlesForStaff, err := models.GetArticlesByCategory("персоналу")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	data := struct {
		LatestArticles []models.Article
	}{
		LatestArticles: articlesForStaff,
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
	}
}

func forApplicantsFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/applicants_path" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/layout_applicants.tmpl",
		"./ui/html/layout_base.tmpl",
		"./ui/html/powered_by.tmpl",
	}

	articlesForApplicants, err := models.GetArticlesByCategory("абитуриентам")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	data := struct {
		LatestArticles []models.Article
	}{
		LatestArticles: articlesForApplicants,
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
	}
}

func forResearchersFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/researchers_path" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/layout_researchers.tmpl",
		"./ui/html/layout_base.tmpl",
		"./ui/html/powered_by.tmpl",
	}

	articlesForResearchers, err := models.GetArticlesByCategory("научные")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	data := struct {
		LatestArticles []models.Article
	}{
		LatestArticles: articlesForResearchers,
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Ошибка сервера", 500)
	}
}

func addArticleFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add_article" {
		http.NotFound(w, r)
		return
	}

	var successMessage, errorMessage string

	switch r.Method {
	case http.MethodGet:
		files := []string{
			"./ui/html/layout_addNew.tmpl",
			"./ui/html/layout_base.tmpl",
			"./ui/html/powered_by.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		title := r.FormValue("title")
		content := r.FormValue("content")
		rawAudience := r.FormValue("audience")

		audience := strings.ToLower(strings.TrimSpace(rawAudience))

		validAudiences := map[string]bool{
			"студентам":    true,
			"персоналу":    true,
			"абитуриентам": true,
			"научные":      true,
		}

		if !validAudiences[audience] {
			errorMessage = "Целевая аудитория указана неверно!!!"
		} else {
			newArticle := models.Article{
				Title:    title,
				Content:  content,
				Audience: audience,
			}

			err := models.AddArticle(newArticle)
			if err != nil {
				log.Println(err.Error())
				errorMessage = "Статья не добавлена, возникла ошибка!!!"
			} else {
				successMessage = "Статья успешно добавлена!!!"
			}
		}

		files := []string{
			"./ui/html/layout_addResult.tmpl",
			"./ui/html/layout_base.tmpl",
			"./ui/html/powered_by.tmpl",
		}

		data := struct {
			SuccessMessage string
			ErrorMessage   string
		}{
			SuccessMessage: successMessage,
			ErrorMessage:   errorMessage,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

		err = ts.Execute(w, data)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
	}
}
