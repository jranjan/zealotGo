package main

import (
    "net/http"
    "io/ioutil"
    "html/template"
)

func main(){
    deploy_service()
}

type Page struct {
    Title string
    Body []byte
}

func deploy_service(){
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    http.ListenAndServe(":2345", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/view/"):]
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/" + title, http.StatusFound)
        return
    }
    renderTemplaate(w, "view", p)
}


func editHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/" + title, http.StatusFound)
        return
    }
    renderTemplaate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    p.save()
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}


func (p *Page) save() error{
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"

    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title:title, Body:body}, nil
}

func renderTemplaate(w http.ResponseWriter, tmpl string, p * Page){
    t, _ := template.ParseFiles(tmpl+".html")
    t.Execute(w, p)
}

