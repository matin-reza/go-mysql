package main

import (
	"fmt"
	"go-mysql/dao"
	"go-mysql/entity"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("template/index.html"))

type Data struct {
	Persons    []entity.Person
	PersonEdit entity.Person
}

var data = Data{}

func Index(w http.ResponseWriter, r *http.Request) {
	data.Persons = dao.ReadAll()

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println(err.Error())
	}
}
func Add(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	person := entity.Person{}
	if id != 0 {
		person.Id = id
	}
	person.Name = r.FormValue("name")
	person.Email = r.FormValue("email")
	dao.Save(person)
	data.PersonEdit = entity.Person{}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	person := dao.LoadById(id)
	data.PersonEdit = person
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	dao.Delete(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/add", Add)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/delete", Delete)
	fmt.Println("Application is started and ready now ...")
	http.ListenAndServe(":8080", nil)
}
