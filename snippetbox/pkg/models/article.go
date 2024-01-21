package models

import "bolatdauren.net/snippetbox/pkg/drivers"

type Article struct {
	ID       int
	Title    string
	Content  string
	Audience string
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
