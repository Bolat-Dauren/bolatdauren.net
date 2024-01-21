package main

import (
	"bolatdauren.net/snippetbox/pkg/drivers"
	"log"
	"net/http"
)

func main() {

	drivers.InitDB("user=postgres dbname=news_db password=0000 sslmode=disable")

	mux := http.NewServeMux()
	mux.HandleFunc("/", homePageFunction)
	mux.HandleFunc("/students_path", forStudentsFunction)
	mux.HandleFunc("/staff_path", forStaffFunction)
	mux.HandleFunc("/applicants_path", forApplicantsFunction)
	mux.HandleFunc("/researchers_path", forResearchersFunction)
	mux.HandleFunc("/add_article", addArticleFunction)
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Println("Starting server on :5000")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
