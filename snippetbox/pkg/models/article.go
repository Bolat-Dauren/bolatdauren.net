// bolatdauren.net/snippetbox/pkg/models/db.go

package models

import (
	"bolatdauren.net/snippetbox/pkg/drivers"
	"database/sql"
	"errors"
	_ "errors"
)

type Article struct {
	ID       int
	Title    string
	Content  string
	Audience string
}

type Meal struct {
	MealName string
	Weekday  string
	Quantity int
}

type User struct {
	ID             string
	FullName       string
	Email          string
	HashedPassword string
	Role           string
}

func RegisterUser(user User) error {
	_, err := drivers.DB.Exec("INSERT INTO users (id, full_name, email, hashed_password, role) VALUES ($1, $2, $3, $4, $5)",
		user.ID, user.FullName, user.Email, user.HashedPassword, user.Role)
	if err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(email, hashedPassword string) (User, error) {
	var user User
	err := drivers.DB.QueryRow("SELECT id, full_name, email, hashed_password, role FROM users WHERE email = $1 AND hashed_password = $2",
		email, hashedPassword).Scan(&user.ID, &user.FullName, &user.Email, &user.HashedPassword, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, errors.New("Пользователь с указанными учетными данными не найден")
		}
		return User{}, err
	}

	return user, nil
}

func AddMeal(meal Meal) error {
	_, err := drivers.DB.Exec("INSERT INTO canteen_menu (meal_name, weekday, quantity) VALUES ($1, $2, $3)",
		meal.MealName, meal.Weekday, meal.Quantity)
	return err
}

func AddArticle(article Article) error {
	_, err := drivers.DB.Exec("INSERT INTO articles (title, content, audience) VALUES ($1, $2, $3)",
		article.Title, article.Content, article.Audience)
	return err
}

func GetLatestArticles() ([]Article, error) {
	rows, err := drivers.DB.Query("SELECT id, title, content FROM articles ORDER BY id DESC LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article

	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func GetArticlesByCategory(category string) ([]Article, error) {
	rows, err := drivers.DB.Query("SELECT title, content FROM articles WHERE audience = $1", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article

	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.Title, &article.Content); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
