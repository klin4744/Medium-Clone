package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"

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
	Topic              string
	Authorimgurl       string
	Authorname         string
	Authoremail        string
	Organizationimgurl string
	Orgname            string
}

func handleArticleRoute(w http.ResponseWriter, req *http.Request) {
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
							  article.claps, article.topic, author.imgurl as authorImgURL, author.name as authorName,
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
			err := rows.Scan(&a.Id, &a.Title, &a.Userid, &a.Content, &a.Organizationid, &a.Articleimgurl, &a.Dateposted, &a.Claps, &a.Topic, &a.Authorimgurl, &a.Authorname, &a.Authoremail, &a.Organizationimgurl, &a.Orgname)
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

func handleSingleArticle(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalln(err)
		}
		defer db.Close()
		vars := mux.Vars(req)
		uri := vars["id"]
		var sqlStatement string
		sqlStatement = `
		SELECT article.id, article.title, article.userid, article.content,
		article.organizationid, article.imgurl as articleImgURL, article.dateposted,
		article.claps, article.topic, author.imgurl as authorImgURL, author.name as authorName,
		author.emailaddress as authorEmail, organization.imgurl as organizationImgURL,
		organization.name as orgname
		FROM  article
		FULL JOIN  organization
			ON article.organizationid = organization.id
		FULL JOIN author
			ON article.userid = author.id
		WHERE article.id = $1
`
		rows, err := db.Query(sqlStatement, uri)
		if uri == "top" {
			sqlStatement = `
							  SELECT article.id, article.title, article.userid, article.content,
							  article.organizationid, article.imgurl as articleImgURL, article.dateposted,
							  article.claps, article.topic, author.imgurl as authorImgURL, author.name as authorName,
							  author.emailaddress as authorEmail, organization.imgurl as organizationImgURL,
							  organization.name as orgname
							  FROM  article
							  FULL JOIN  organization
							  	ON article.organizationid = organization.id
							  FULL JOIN author
								ON article.userid = author.id
							  ORDER BY
							  	article.claps DESC
							  LIMIT 4

		`
			rows, err = db.Query(sqlStatement)
		}
		defer rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
		var articles []*Article
		for rows.Next() {
			a := new(Article)
			err := rows.Scan(&a.Id, &a.Title, &a.Userid, &a.Content, &a.Organizationid, &a.Articleimgurl, &a.Dateposted, &a.Claps, &a.Topic, &a.Authorimgurl, &a.Authorname, &a.Authoremail, &a.Organizationimgurl, &a.Orgname)
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
	}

}

type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	// http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	// http.Handle("/articles/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	// http.Handle("/api/", rtr)

	router := mux.NewRouter()

	router.HandleFunc("/api/articles", handleArticleRoute)
	router.HandleFunc("/api/articles/{id:.+}", handleSingleArticle)

	spa := spaHandler{staticPath: "./static", indexPath: "./static/index.html"}
	router.PathPrefix("/articles/").Handler(spa)
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
