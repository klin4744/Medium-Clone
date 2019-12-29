package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var connStr string = "user=kevin2 password=1234 dbname=medium sslmode=disable"

// All things meant to be exported must be capitalized, including to any writer: IE exported to browser
type Article struct {
	Id                 int
	Title              string
	Userid             int
	Content            string
	Organizationid     int
	Articleimgurl      string
	Dateposted         string
	Claps              int
	Authorimgurl       string
	Authorname         string
	Authoremail        string
	Organizationimgurl string
	Orgname            string
}

func handleArticle(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	switch req.Method {
	case "GET":
		var articles []*Article
		sqlStatement := `SELECT article.id, article.title, article.userid, article.content,
							  article.organizationid, article.imgurl as articleImgURL, article.dateposted,
							  article.claps, author.imgurl as authorImgURL, author.name as authorName,
							  author.emailaddress as authorEmail, organization.imgurl as organizationImgURL,
							  organization.name as orgname
							  FROM  article
							  FULL JOIN  organization
							  	ON article.organizationid = organization.id
							  FULL JOIN author
							  	ON article.userid = author.id
								`
		rows, err := db.Query(sqlStatement)
		defer rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
		for rows.Next() {
			a := new(Article)
			err := rows.Scan(&a.Id, &a.Title, &a.Userid, &a.Content, &a.Organizationid, &a.Articleimgurl, &a.Dateposted, &a.Claps, &a.Authorimgurl, &a.Authorname, &a.Authoremail, &a.Organizationimgurl, &a.Orgname)
			if err != nil {
				log.Fatalln(err)
			}
			articles = append(articles, a)
		}
		fmt.Println("200 SUCCESS")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(articles)
		if err != nil {
			log.Fatalln(err)
		}
		return
	case "POST":
		fmt.Fprintf(w, "POST")
	default:
		log.Fatalln("Bad request, route does not exist")
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/articles", handleArticle)
	http.ListenAndServe(":1000", nil)
}
