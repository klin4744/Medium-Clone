package main

import (
	"fmt"
	"log"
	"net/http"
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
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/article", handleArticle)
	http.ListenAndServe(":3111", nil)
}
