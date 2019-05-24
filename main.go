package main

import (
	"blogg/models"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

var posts map[string]*models.Post

//==================================================================PARSFILE================================
var tpl *template.Template

func init() {

	var err error
	tpl, err = template.ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/write.html",
	)

	if err != nil {
		panic(err)
	}
}

//==================================================================INDEX================================
func IndexH(rw http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(rw, "index", posts)
	if err != nil {
		fmt.Printf("ExecuteTemplate INDEX: %s\n", err)
	}

}

//==================================================================WRITE================================
func WriteH(rw http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(rw, "write", nil)
	if err != nil {
		fmt.Printf("ExecuteTemplate WRITE: %s\n", err)
	}
}

//==================================================================SAVEPOST================================
func SavePostH(rw http.ResponseWriter, r *http.Request) {

	id := strconv.Itoa(rand.Intn(31))
	title := r.FormValue("title")
	content := r.FormValue("content")
	fmt.Printf("%s\n", id)

	post := models.NewPost(id, title, content)

	posts[post.ID] = post

	http.Redirect(rw, r, "/", 302)
}

func main() {

	posts = make(map[string]*models.Post, 0)

	//для подключения стилей        убираем префикс ""
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	http.HandleFunc("/", IndexH)

	http.HandleFunc("/write", WriteH)

	http.HandleFunc("/SavePost", SavePostH)

	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		panic(err)
	}

}
