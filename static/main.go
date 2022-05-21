package main

import (
	"log"
	"database/sql"
   "html/template"
	_ "github.com/go-sql-driver/mysql"
    "net/http"
    //"io/ioutil"
)

type Users struct{
    Pseudo string
    Adresse_mail string
    Password string
}

const (
	Host = "localhost"
	Port = "8080"
)

var tmpl = template.Must(template.ParseFiles("static/index.html"))
var filesserver = http.FileServer(http.Dir("static/css/"))

func pages(w http.ResponseWriter, r *http.Request){
    register : Users {
        Pseudo : r.FormValue("pseudo"),
        Adresse_mail : r.FormValue("adresse_mail"),
        Password : r.FormValue("password"),
    }
    
    tmpl.Execute(w, tmpl)
}
func dbz(){
    db, err := sql.Open( "mysql", "root:@tcp(localhost:3306)/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    rows, err:= db.Query("INSERT INTO users (`id`, `pseudo`, `adresse_mail`, `password`) VALUES (DEFAULT,'Guigui','guillaume@gmail.com','Guillaume2004')")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    
    /*for rows.Next() {
        var id_user int
        var firstname string
        var lastname string
        var password string
        if err := rows.Scan(&id_user, &firstname, &lastname, &password); err != nil {
            log.Fatal(err)
        }
        fmt.Println(id_user, firstname, lastname, password)
    }*/
}
func main() {
    print("Lancement de la page instanciée sur : " + Host + ":" + Port )
    
    http.HandleFunc("/",pages)
    http.ListenAndServe(Host+":"+Port, nil)
    http.Handle("/css/", http.StripPrefix("/css/", filesserver))
}
