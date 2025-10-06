package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
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
type UpdateBlog struct {
	Title      string `json:"title"`
	NewTitle   string `json:"new_title"`
	NewContent string `json:"new_content"`
}

var mutex sync.RWMutex

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:example@tcp(172.35.0.3:3306)/blog")
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

func CheckEmpty(post Blog, update UpdateBlog) error {
	err := errors.New("empty set")

	if !reflect.ValueOf(post).IsZero() {
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
	if !reflect.ValueOf(update).IsZero() {
		switch "" {
		case update.Title:
			fmt.Println("Title cannot be empty")
			return err
		case update.NewTitle:
			fmt.Println("Content cannot be empty")
			return err
		case update.NewContent:
			fmt.Println("Category cannot be empty")
			return err
		default:
			return nil

		}
	}
	return nil

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
	j, err := json.Marshal(Bl)
	CheckError(err)
	w.Write(j)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Blog
	var up UpdateBlog
	db, err := connectDB()
	CheckError(err)

	err = json.NewDecoder(r.Body).Decode(&post)
	CheckError(err)
	err = CheckEmpty(post, up)

	CheckError(err)
	query := "INSERT INTO `post` (`title`,`content`,`category`,`tags`) VALUES(?,?,?,?)"
	mutex.Lock()
	res, err := db.ExecContext(context.Background(), query, post.Title, post.Content, post.Category, post.Tags)
	mutex.Unlock()
	CheckError(err)
	fmt.Printf("response: %s ", res)
}

func CheckExists(db *sql.DB, title string) (error, *sql.Rows) {
	err1 := errors.New("Not found")
	var temp string

	query := "select title from post where title=?;"
	mutex.Lock()
	err := db.QueryRow(query, title).Scan(&temp)
	fmt.Println("Result is ", temp)
	mutex.Unlock()
	CheckError(err)

	fmt.Printf("\nTHE TEST response: %s \n", title)
	fmt.Printf("\nTHE TEMP IS %s\n", temp)
	if temp != title {
		fmt.Println("IN IF WHERE NO QUERY IS DONE ")
		return err1, nil
	} else {

		Result, err := db.Query("select * from post where title=?;", title)
		CheckError(err)
		return nil, Result
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	CheckError(err)
	var post UpdateBlog
	var bl Blog
	err = json.NewDecoder(r.Body).Decode(&post)
	CheckError(err)
	err = CheckEmpty(bl, post)
	CheckError(err)
	err, _ = CheckExists(db, post.Title)
	CheckError(err)

	query := "UPDATE post SET title = ?, content = ? WHERE title = ?;"

	mutex.Lock()
	res2, err := db.ExecContext(context.Background(), query, post.NewTitle, post.NewContent, post.Title)
	mutex.Unlock()
	CheckError(err)
	fmt.Printf("response: %s ", res2)
}

func SearchPost(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	var JSON_post Blog
	CheckError(err)
	id := string(r.PathValue("id"))

	fmt.Printf("Id is %s", id)
	err, post := CheckExists(db, id)
	CheckError(err)
	for post.Next() {
		err := post.Scan(&JSON_post.Title, &JSON_post.Content, &JSON_post.Category, &JSON_post.Tags, &JSON_post.Created, &JSON_post.Updated)
		CheckError(err)
	}
	fmt.Println(JSON_post.Title, JSON_post.Content, JSON_post.Category, JSON_post.Tags, JSON_post.Created, JSON_post.Updated)
	j, err := json.Marshal(JSON_post)
	CheckError(err)
	w.Write(j)
	w.WriteHeader(http.StatusOK)
}

// TODO: make delete post feature DELETE FROM table_name WHERE condition;
func DeletePost(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	CheckError(err)
	id := r.PathValue("id")
	err, _ = CheckExists(db, id)
	CheckError(err)
	_, err = db.Query("DELETE FROM post WHERE `title`=?", id)
	CheckError(err)
	w.WriteHeader(http.StatusGone)
}
