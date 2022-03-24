package main

import (
	"reflect"
	"testing"
)

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {

			t.Fail()
			break
		}
	}
}

func Test_getArticleById(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    *article
		wantErr bool
	}{
		{"should find an existing article", 1, &articleList[0], false},
		{"should return an error on non-existent article", 10, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getArticleById(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getArticleById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getArticleById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
