// bolatdauren.net/snippetbox/cmd/web/main/handlers.go

package main

import (
	"bolatdauren.net/snippetbox/pkg/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

func addMealPageFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add_meal" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		files := []string{
			"./ui/html/layout_add_meal.tmpl", // Use the new meal creation template
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
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

		// Extract form values
		mealName := r.Form.Get("meal_name")
		weekday := r.Form.Get("weekday")
		quantityStr := r.Form.Get("quantity")

		// Convert quantity to integer
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

		// Create a Meal struct
		meal := models.Meal{
			MealName: mealName,
			Weekday:  weekday,
			Quantity: quantity,
		}

		// Add meal to database
		err = models.AddMeal(meal)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

		// Redirect to a success page or handle as needed
		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
	}
}

func registerPageFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/layout_register.tmpl",
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
}

func loginPageFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/layout_login.tmpl",
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
}

func registerUserFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

		// Extract form values
		id := r.Form.Get("id")
		fullName := r.Form.Get("full_name")
		email := r.Form.Get("email")
		hashedPassword := r.Form.Get("hashed_password")
		role := r.Form.Get("role")

		// Create a User struct
		user := models.User{
			ID:             id,
			FullName:       fullName,
			Email:          email,
			HashedPassword: hashedPassword,
			Role:           role,
		}

		// Add user to database
		err = models.RegisterUser(user)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/registration_success", http.StatusSeeOther)

	default:
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
	}
}

func loginFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

		// Extract form values
		email := r.Form.Get("email")
		hashedPassword := r.Form.Get("hashed_password")

		// Validate user credentials
		user, err := models.AuthenticateUser(email, hashedPassword)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

		// Redirect based on user role
		switch user.Role {
		case "admin":
			http.Redirect(w, r, "/admin_dashboard", http.StatusSeeOther)
		case "teacher":
			http.Redirect(w, r, "/teacher_dashboard", http.StatusSeeOther)
		case "student":
			http.Redirect(w, r, "/student_dashboard", http.StatusSeeOther)
		default:
			http.Error(w, "Неверная роль пользователя", http.StatusForbidden)
		}

	default:
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
	}
}

// Define other handler functions...

// Define registration success page handler
func registrationSuccessPageFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration_success" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/layout_registration_success.tmpl",
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
}
