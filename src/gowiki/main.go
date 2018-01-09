package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"html/template"
	"os"
)

var RESSOURCE_PATH string

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
	}
	files, err := ioutil.ReadDir(RESSOURCE_PATH+"/data")
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error..."))
		return
	}
	filename := make([]string, len(files))
	for i := range files {
		filename[i] = strings.Replace(files[i].Name(), ".txt", "", 1)
	}
	t, _ := template.ParseFiles(RESSOURCE_PATH+"/templates/index.html")
    t.Execute(w, filename)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/v/"):]
	p,err := load(title);
	if err != nil {
		http.Redirect(w,r,"/n/"+title,302)
		return
	}
	t, _ := template.ParseFiles(RESSOURCE_PATH+"/templates/view.html")
    t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/v/"):]
	p,err := load(title);
	if err != nil {
		http.Redirect(w,r,"/n/"+title,302)
		return
	}
	t, _ := template.ParseFiles(RESSOURCE_PATH+"/templates/edit.html")
    t.Execute(w, p)
}
func newHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/n/"):]
	t, _ := template.ParseFiles(RESSOURCE_PATH+"/templates/new.html")
    t.Execute(w, title)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/s/"):]
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error..."))
		return
	}
	fmt.Println(r.PostForm)
	if title == "" {
		title = r.PostForm["title"][0]
	}
	body := r.PostForm["body"][0]
	(&(Page{title, []byte(body)})).save()
	http.Redirect(w,r,"/v/"+title,302)
}
func main() {
	RESSOURCE_PATH = os.Getenv("GOPATH")+"/ressources/gowiki"
	fmt.Println("Load your gowiki listen on 8080...")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/v/", viewHandler)
	http.HandleFunc("/e/", editHandler)
	http.HandleFunc("/n/", newHandler)
	http.HandleFunc("/s/", saveHandler)
	http.ListenAndServe(":8080",nil)
}
