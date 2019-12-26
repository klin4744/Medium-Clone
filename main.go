package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func handleArticle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprint(w, "GET Request sent")
	case "POST":
		fmt.Fprintf(w, "POST")
	case "DELETE":
		fmt.Fprintf(w, "DELETE")
	case "PUT":
		fmt.Fprintf(w, "PUT")
	default:
		log.Fatalln("An Error Occured")
	}
}

func main() {
	connStr := "user=kevin2 password=1234 dbname=medium sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/article", handleArticle)
	http.ListenAndServe(":3111", nil)
}
