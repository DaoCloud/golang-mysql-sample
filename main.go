package main

import (
	"html/template"
	"log"
	"net/http"
)

func init() {
	Config()
	MustConnectDB()
	InitDB()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/new", insert)
	http.HandleFunc("/drop", drop)

	log.Println("Start listening...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
			res.WriteHeader(http.StatusInternalServerError)
		}
	}()

	log.Println("Index home")

	t, err := template.New("foo").Parse(string(tpl))
	if err != nil {
		log.Println(err)
		res.WriteHeader(500)
		return
	}

	data := make(map[string]interface{})
	data["List"] = List()
	t.Execute(res, data)
}

func insert(res http.ResponseWriter, req *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
			res.WriteHeader(http.StatusInternalServerError)
		}
	}()

	person := &Person{}
	person.Name = req.FormValue("name")
	person.Phone = req.FormValue("phone")

	Insert(person)

	log.Printf("Insert new person %v\n", *person)
	http.Redirect(res, req, "/", 302)
}

func drop(res http.ResponseWriter, req *http.Request) {
	log.Println("drop collection")

	Drop()

	http.Redirect(res, req, "/", 302)
}
