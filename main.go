package main

import (
	"blogg/documents"
	"blogg/models"
	"blogg/session"
	"blogg/utils"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"gopkg.in/mgo.v2" //использ mongo
)

const (
	COOKIE_NAME = "sessionId"
)

var postsCollection *mgo.Collection
var inMemorySession *session.Session

//==================================================================PARSFILE================================
var tpl *template.Template

func init() {

	var err error
	tpl, err = template.ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/write.html",
		"templates/login.html",
	)

	if err != nil {
		panic(err)
	}
}

//==================================================================INDEX================================
func IndexH(rw http.ResponseWriter, r *http.Request) {

	cookie, errr := r.Cookie(COOKIE_NAME)
	if errr != nil {
		fmt.Printf("No COOKIE: %s\n", errr)
	}
	if cookie != nil {
		fmt.Printf("USERNAME: %v\n", inMemorySession.Get(cookie.Value))
	}

	postDocuments := []documents.PostDocuments{}
	postsCollection.Find(nil).All(&postDocuments)

	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{doc.ID, doc.Title, doc.Content}
		posts = append(posts, post)
	}
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
	postDocuments := documents.PostDocuments{}
	err := postsCollection.FindId(id).One(&postDocuments)
	if err != nil {
		http.Redirect(rw, r, "/", 302)
		return
	}

	post := models.Post{postDocuments.ID, postDocuments.Title, postDocuments.Content}

	error := tpl.ExecuteTemplate(rw, "write", post)
	if error != nil {
		fmt.Printf("ExecuteTemplate EDIT: %s\n", error)
	}
} //==================================================================DELETE================================
func DeleteH(rw http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	if id == "" {
		http.NotFound(rw, r)
	}

	postsCollection.RemoveId(id)

	http.Redirect(rw, r, "/", 302)
}

//===================================================================SAVEPOST==============================
func SavePostH(rw http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	fmt.Printf("%s\n", id)

	postDocuments := documents.PostDocuments{id, title, content}
	if id != "" { //проверка на наличие поста
		postsCollection.UpdateId(id, postDocuments)
	} else { //создаем новый пост
		id := utils.GenerateId()
		postDocuments.ID = id
		postsCollection.Insert(postDocuments)
	}

	http.Redirect(rw, r, "/", 302)
}

//===================================================================LOGIN==============================

func LoginH(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		err := tpl.ExecuteTemplate(rw, "login", nil)
		if err != nil {
			fmt.Printf("ExecuteTemplate LOGIN: %s\n", err)
		}
	}
	if r.Method == "POST" {

		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Printf("Reg: UserName: %v\n    Password: %v ", username, password)

		sessionId := inMemorySession.Init(username)

		cookie := &http.Cookie{
			Name:    COOKIE_NAME,
			Value:   sessionId,
			Expires: time.Now().Add(5 * time.Minute),
		}

		http.SetCookie(rw, cookie)

		http.Redirect(rw, r, "/", 302)
	}

}

func main() {

	inMemorySession = session.NewSession()

	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Printf("SESSION MG: %s\n", err)
	}

	postsCollection = session.DB("blog").C("posts")

	//для подключения стилей        убираем префикс ""
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	http.HandleFunc("/", IndexH)

	http.HandleFunc("/write", WriteH)

	http.HandleFunc("/edit", EditH)

	http.HandleFunc("/delete", DeleteH)

	http.HandleFunc("/SavePost", SavePostH)

	http.HandleFunc("/login", LoginH)

	error := http.ListenAndServe(":3030", nil)
	if err != nil {
		fmt.Printf("ListenAndServe: %s\n", error)
	}
}
