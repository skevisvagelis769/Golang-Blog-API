package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Blog struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Tags     string `json:"tags"`
	Created  string `json:"created_at"`
	Updated  string `json:"updated_at"`
}

var mutex sync.RWMutex

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:8716@tcp(localhost:3306)/blog")
	CheckError(err)

	err = db.Ping()
	CheckError(err)

	return db, nil

}

func CheckError(err error) error {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return err

}

func CheckEmpty(post Blog) error {
	err := errors.New("empty set")
	switch "" {
	case post.Title:
		fmt.Println("Title cannot be empty")
		return err
	case post.Content:
		fmt.Println("Content cannot be empty")
		return err
	case post.Category:
		fmt.Println("Category cannot be empty")
		return err
	case string(post.Tags):
		fmt.Println("Tags cannot be empty")
		return err
	default:
		return nil

	}
}

func printDb(tb Blog, test []byte) {
	fmt.Println("===========NORMAL STRING==========")
	fmt.Println(tb.Title + " | " + tb.Content + " | " + tb.Category + " | " + string(tb.Tags) + " | " + tb.Created + " | " + tb.Updated)
	fmt.Println("==========IN JSON==========")
	fmt.Println(string(test))
}

func GetALlPost(w http.ResponseWriter, r *http.Request) {

	Bl := []Blog{}
	fmt.Println("this is a test!")
	db, err := connectDB()
	CheckError(err)

	defer db.Close()

	rows, err := db.Query("select * from post;")
	CheckError(err)

	defer rows.Close()

	for rows.Next() {
		var B Blog
		err = rows.Scan(&B.Title, &B.Content, &B.Category, &B.Tags, &B.Created, &B.Updated)
		CheckError(err)

		Bl = append(Bl, B)
	}

	for _, tb := range Bl {
		j, err := json.Marshal(tb)
		CheckError(err)

		printDb(tb, j)
	}
}

// TODO: add mutex when accessing the db to write
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Blog
	db, err := connectDB()
	CheckError(err)

	err = json.NewDecoder(r.Body).Decode(&post)
	CheckError(err)
	err = CheckEmpty(post)

	CheckError(err)
	query := "INSERT INTO `post` (`title`,`content`,`category`,`tags`) VALUES(?,?,?,?)"
	res, err := db.ExecContext(context.Background(), query, post.Title, post.Content, post.Category, post.Tags)
	CheckError(err)
	fmt.Printf("response: %s ", res)
}
