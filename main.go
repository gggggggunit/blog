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

//==================================================================EDIT================================
func EditH(rw http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	post, found := posts[id]
	if !found {
		http.NotFound(rw, r)
	}

	err := tpl.ExecuteTemplate(rw, "write", post)
	if err != nil {
		fmt.Printf("ExecuteTemplate EDIT: %s\n", err)
	}
} //==================================================================DELETE================================
func DeleteH(rw http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	if id == "" {
		http.NotFound(rw, r)
	}

	delete(posts, id)

	http.Redirect(rw, r, "/", 302)
}

//===================================================================SAVEPOST==============================
func SavePostH(rw http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	fmt.Printf("%s\n", id)

	var post *models.Post
	if id != "" { //проверка на наличие поста
		post = posts[id]
		post.Title = title
		post.Content = content
	} else { //создаем новый пост
		id := strconv.Itoa(rand.Intn(31))
		post := models.NewPost(id, title, content)
		posts[post.ID] = post
	}

	http.Redirect(rw, r, "/", 302)
}

func main() {

	posts = make(map[string]*models.Post, 0)

	//для подключения стилей        убираем префикс ""
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	http.HandleFunc("/", IndexH)

	http.HandleFunc("/write", WriteH)

	http.HandleFunc("/edit", EditH)

	http.HandleFunc("/delete", DeleteH)

	http.HandleFunc("/SavePost", SavePostH)

	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		panic(err)
	}
}
