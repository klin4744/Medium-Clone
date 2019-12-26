package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=kevin2 password=1234 dbname=medium sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	stmt, err := db.Prepare(`
	CREATE TABLE Author (
		id  SERIAL    PRIMARY KEY      NOT NULL,
		name               TEXT        NOT NULL,
		emailAddress    VARCHAR(200)   NOT NULL,
		imgUrl          VARCHAR(200)
		)`)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Table created successfully..")
	}
	stmt, err = db.Prepare(`
	CREATE TABLE Article(
		id      SERIAL PRIMARY KEY  NOT NULL,
		title          VARCHAR(200) NOT NULL,
		userId         INT          NOT NULL,
		content        TEXT         NOT NULL,
		organizationId INT,
		imgUrl         VARCHAR(200) NOT NULL,
		datePosted     VARCHAR(200),
		claps          INT
	)
	`)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Table created successfully..")
	}
	stmt, err = db.Prepare(`
	CREATE TABLE AuthorAuthor (
		id      SERIAL PRIMARY KEY NOT NULL ,
		user1Id        INT         NOT NULL,
		user2Id        INT         NOT NULL
	)
	`)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Table created successfully..")
	}
	stmt, err = db.Prepare(`
	CREATE Table Organization(
		id SERIAL PRIMARY KEY NOT NULL,
		imgUrl   VARCHAR(200) NOT NULL,
		pageUrl  VARCHAR(200) NOT NULL
	)
	`)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Table created successfully..")
	}
	defer db.Close()
}
