package controllers

import (
	"github.com/go-chi/chi"
	"net/http"
	"obas/src/config"
	io "obas/src/io/institutions"
)

func Universitys(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", universitysHandler(app))
	return r
}

func universitysHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allinst, err := io.GetUniversity()

		if err != nil {
			app.ServerError(w, err)
		}

		type PageData struct {
			universitys []io.Universitys
			name        string
		}
		data := PageData{allinst, ""}

		files := []string{
			app.Path + "",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, "", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}