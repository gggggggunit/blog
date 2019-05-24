package handlers

import "html/template"

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
