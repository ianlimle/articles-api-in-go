// Create a REST API that allows us to CREATE, READ, UPDATE and DELETE the articles on our website.
// When we talk about CRUD APIs we are referring to an API that can handle all of these tasks:
// Creating, Reading, Updating and Deleting.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// define Article structure
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// declare a global Articles array
// that we can then populate in our main function to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)
	// unmarshall this into a new Article struct
	// append this to our Articles array
	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include new Article
	Articles = append(Articles, article)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	fmt.Println("Endpoint Hit: createNewArticle")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	// encode articles array into a JSON string and write as part of response
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	// obtain {id} value from our URL and return the article that matches this criteria
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "key: "+key)

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == key {
			// encode article into a JSON string and write as part of response
			json.NewEncoder(w).Encode(article)
		}
	}
	fmt.Println("Endpoint Hit: returnSingleArticle")
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedArticle Article
	json.Unmarshal(reqBody, &updatedArticle)

	// loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our articles
		if article.Id == id {
			// update our Articles array to update the article
			Articles[index] = updatedArticle
		}
	}
	fmt.Println("Endpoint Hit: updateArticle")
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// parse the path parameters
	vars := mux.Vars(r)
	// extract the `id` of the article to delete
	id := vars["id"]

	// loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our articles
		if article.Id == id {
			// update our Articles array to remove the article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
	fmt.Println("Endpoint Hit: deleteArticle")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	// add 'all' route and map it to our returnAllArticles function
	myRouter.HandleFunc("/all", returnAllArticles)
	// add 'article' route and map it to our createNewArticle function as a 'HTTP POST' request
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	// add 'article/{id}' route and map it to our deleteArticle function as a 'HTTP DELETE' request
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	// add 'article/{id}' route and map it to our updateArticle function as a 'HTTP PUT' request
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	// add 'article/{id}' route and map it to our returnSingleArticle function
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
