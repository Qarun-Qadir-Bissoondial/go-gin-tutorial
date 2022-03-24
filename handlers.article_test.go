package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

func TestArticleListRendering(t *testing.T) {
	encoding := "application/json"
	r := getRouter(false)
	r.GET("/", showIndexPage)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", encoding)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		responseTypeOK := strings.Contains(w.Result().Header.Get("Content-Type"), encoding)
		expectedBytes, _ := json.Marshal(articleList)
		bodyOK := string(p) == string(expectedBytes)

		if !statusOK {
			t.Errorf("Expected %d, got %d\n", http.StatusOK, w.Code)
		}

		if !responseTypeOK {
			t.Errorf("Expected %s, got %s\n", encoding, w.Result().Header.Get("Content-Type"))
		}

		if !bodyOK {
			t.Errorf("Expected %s, got %s\n", string(expectedBytes), string(p))
		}

		return err == nil && statusOK && responseTypeOK && bodyOK
	})

	encoding = "application/xml"
	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", encoding)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		responseTypeOK := strings.Contains(w.Result().Header.Get("Content-Type"), encoding)
		expectedBytes, _ := xml.Marshal(articleList)
		bodyOK := string(p) == string(expectedBytes)

		if !statusOK {
			t.Errorf("Expected %d, got %d\n", http.StatusOK, w.Code)
		}

		if !responseTypeOK {
			t.Errorf("Expected %s, got %s\n", encoding, w.Result().Header.Get("Content-Type"))
		}

		if !bodyOK {
			t.Errorf("Expected %s, got %s\n", string(expectedBytes), string(p))
		}

		return err == nil && statusOK && responseTypeOK && bodyOK
	})
}

func TestSingleArticleRendering(t *testing.T) {
	encoding := "text/html"
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)
	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		responseTypeOK := strings.Contains(w.Result().Header.Get("Content-Type"), encoding)
		p, err := ioutil.ReadAll(w.Body)
		bodyOK := strings.Contains(string(p), "<h1>Article 1</h1>")

		if !statusOK {
			t.Errorf("Expected %d, got %d\n", http.StatusOK, w.Code)
		}

		if !responseTypeOK {
			t.Errorf("Expected %s, got %s\n", encoding, w.Result().Header.Get("Content-Type"))
		}

		return err == nil && statusOK && responseTypeOK && bodyOK
	})

	req, _ = http.NewRequest("GET", "/article/view/55", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusNotFound
		return statusOK
	})

	req, _ = http.NewRequest("GET", "/article/view/invalid_id", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == http.StatusBadRequest
	})
}

func TestArticleCreationAuthenticated(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()

	r := getRouter(true)

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	r.POST("/article/create", createArticle)

	articlePayload := getArticlePOSTPayload()
	req, _ := http.NewRequest("POST", "/article/create", strings.NewReader(articlePayload))
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(articlePayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Submission Successful</title>") < 0 {
		t.Fail()
	}
	restoreLists()
}

func getArticlePOSTPayload() string {
	params := url.Values{}
	params.Add("title", "Test Article Title")
	params.Add("content", "Test Article Content")

	return params.Encode()
}
