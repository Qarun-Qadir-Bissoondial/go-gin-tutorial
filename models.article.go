package main

import (
	"errors"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

//getAllArticles returns a list of all articles in memory
func getAllArticles() []article {
	return articleList
}

func getArticleById(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("article not found")
}

func createNewArticle(title, content string) (*article, error) {
	a := article{ID: len(articleList) + 1, Title: title, Content: content}

	articleList = append(articleList, a)

	return &a, nil
}
