package main

import (
	"log"
	"database/sql"
   "html/template"
	_ "github.com/go-sql-driver/mysql"
    "net/http"
    "fmt"
    
    //"io/ioutil"
)

type Users struct{
    Pseudo string
    Adresse_mail string
    Password string
}
var register Users

const (
	Host = "localhost"
	Port = "8080"
)



var tmpl = template.Must(template.ParseFiles("static/index.html"))
var filesserver = http.FileServer(http.Dir("static/css/"))

func pages(w http.ResponseWriter, r *http.Request){
    register = Users {
        Pseudo : r.FormValue("pseudo"),
        Adresse_mail : r.FormValue("adresse_mail"),
        Password : r.FormValue("password"),
    }
    dbz(w, r)
    tmpl.Execute(w, tmpl)


}
func dbz(w http.ResponseWriter, r *http.Request){
    db, err := sql.Open( "mysql", "root:@tcp(localhost:3306)/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    rows, err:= db.Query("INSERT INTO users (`id`, `pseudo`, `adresse_mail`, `password`) VALUES (DEFAULT,?,?,?)", register.Pseudo, register.Adresse_mail, register.Password)
    if err != nil {
        log.Fatal(err)
    }
    query, err:= db.Query("SELECT pseudo FROM users where pseudo = " + register.Pseudo)
    bouh := ""
    query.Scan(&bouh)
    var oui = "<p>bouh</p>"
    fmt.Fprint(w ,oui)
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
func search(){
    /*bouh, err:= db.Query("SElECT FROM users (`pseudo`) where 'pseudo'="+register.Pseudo).Scan()
    const oui = "<p>"+bouh+"</p>"
    fmt.Fprint(w, oui)*/
    
}
func main() {
    print("Lancement de la page instanciée sur : " + Host + ":" + Port ) 
    http.HandleFunc("/test", pages)
    
    http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("static"))))
    //http.Handle("/css", http.StripPrefix("/css", filesserver))
    http.Handle("/", http.FileServer(http.Dir("css/")))
    http.ListenAndServe(Host+":"+Port, nil)
    
}
