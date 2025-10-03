package handlers

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

/* func printDb(tb Blog, test []byte) {
	fmt.Println("===========NORMAL STRING==========")
	fmt.Println(tb.Title + " | " + tb.Content + " | " + tb.Category + " | " + string(tb.Tags) + " | " + tb.Created + " | " + tb.Updated)
	fmt.Println("==========IN JSON==========")
	fmt.Println(string(test))
} */

func printDb(tb Blog, test []byte) {
	fmt.Println("===========NORMAL STRING==========")
	fmt.Println(tb.Title + " | " + tb.Content + " | " + tb.Category + " | " + string(tb.Tags) + " | " + tb.Created + " | " + tb.Updated)
	//send, err := json.Marshal(tb)
	/* 	var jsonstr = fmt.Sprintf(``)
	 */
	data, err := os.ReadFile("/media/carabi/New Volume/GO/src/github.com/skevisvagelis769/BLOG_API/handlers/test.json")
	CheckError(err)
	//send, err := json.Marshal(data)
	CheckError(err)
	req, err := http.NewRequest("POST", "http://localhost:4318/v1/logs", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	CheckError(err)
	client := &http.Client{}
	resp, err := client.Do(req)
	CheckError(err)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// 3. Print response
	fmt.Println("==========HTTP RESPONSE==========")
	fmt.Println("Status:", resp.Status)
	fmt.Println("Body:", string(body))
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

//TODO: make an update post feature (update post set title = 'CHANGED UP' , content = 'THIS IS A CHANGED MSG' where title ='Wawaweewa';)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	CheckError(err)
	var post UpdateBlog
	var bl Blog
	err = json.NewDecoder(r.Body).Decode(&post)
	CheckError(err)
	err = CheckEmpty(bl, post)
	CheckError(err)
	query := "UPDATE post SET title = ?, content = ? WHERE title = ?;"
	mutex.Lock()
	res, err := db.ExecContext(context.Background(), query, post.NewTitle, post.NewContent, post.Title)
	mutex.Unlock()
	CheckError(err)
	fmt.Printf("response: %s ", res)
}

//TODO: make search post by title feature

//TODO: make delete post feature
