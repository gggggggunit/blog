package main

import (
	"blogg/handlers"
	"net/http"
)

func main() {

	//для подключения стилей        убираем префикс ""
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	http.HandleFunc("/", handlers.IndexH)
	http.HandleFunc("/write", handlers.WriteH)
	http.HandleFunc("/SavePost", handlers.SavePostH)

	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		panic(err)
	}

}
