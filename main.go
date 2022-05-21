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
    id int
    Pseudo string
    Adresse_mail string
    Password string
}
type ALL struct{
    User Users
    connect bool
}

const (
	Host = "localhost"
	Port = "8070"
)

var data = ALL{}
var db *sql.DB
var tmpl *template.Template
var filesserver = http.FileServer(http.Dir("static/css/"))

func pages(w http.ResponseWriter, r *http.Request){
    if r.FormValue("submit") != "" {
        register := Users {
            Pseudo : r.FormValue("pseudo"),
            Adresse_mail : r.FormValue("adresse_mail"),
            Password : r.FormValue("password"),
        }
        dbz(w, r, register)

    }
    tmpl.ExecuteTemplate(w, "home", nil)


}
/*func user(){
    rows, err:= db.Query(fmt.Sprintf("INSERT INTO users (`id`, `pseudo`, `adresse_mail`, `password`) VALUES (DEFAULT,'%s', '%s', '%s')", register.Pseudo, register.Adresse_mail, register.Password))
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
}*/
func connectUser(name string) (Users, error){
    user := Users{}
    getuser:= fmt.Sprintf("SELECT * FROM users WHERE pseudo='%s'", name)
    err := db.QueryRow(getuser).Scan(&user.id, &user.Pseudo,&user.Adresse_mail, &user.Password)
    return user, err
}
func login(w http.ResponseWriter, r *http.Request){
    log.Println("funtion calling")
    user := Users{}
    name := r.FormValue("username")
    password := r.FormValue("password")
    log.Println("test")
    
    if r.FormValue("submit") == "connect"{
        log.Println("oui")
        connectUser(name)
        log.Println(password)
        log.Println(user.Password)
        /*if password == user.Password{
            log.Println(user.Pseudo)
            data.connect = true
            //http.Redirect(w, r, "http://" + Host + ":" + Port + "/test", http.StatusMovedPermanently)
        }else{
            log.Println("wrong password")
            data.connect = false
        }*/
    }
    tmpl.ExecuteTemplate(w, "account", data)
}
func dbz(w http.ResponseWriter, r *http.Request, register Users){
    rows, err:= db.Query("INSERT INTO users (`pseudo`, `adresse_mail`, `motdepasse`) VALUES (?,?,?)", register.Pseudo, register.Adresse_mail, register.Password)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(fmt.Sprintf("SELECT pseudo FROM users where pseudo = '%s'", register.Pseudo))
    var bouh string
    db.QueryRow(fmt.Sprintf("SELECT `pseudo` FROM `users` where `pseudo` = '%s'", register.Pseudo)).Scan(&bouh)
    var oui = fmt.Sprintf("<p>%s</p>", bouh)
    fmt.Fprint(w, oui)
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
    db, _ = sql.Open( "mysql", "root:@tcp(localhost:3306)/testdb")
    defer db.Close()

    pageServer := http.FileServer(http.Dir("static/html"))
    http.Handle("/html/", http.StripPrefix("/html/", pageServer) )
    styleServer := http.FileServer(http.Dir("static/css"))
    http.Handle("/css/", http.StripPrefix("/css/", styleServer) )

    var err error
    tmpl, err = template.New("").ParseGlob("static/html/*.html")
    log.Print(err)

    http.HandleFunc("/test", pages)
    http.HandleFunc("/connect", login)
    
    print("Lancement de la page instanciée sur : " + Host + ":" + Port ) 
    http.ListenAndServe(Host+":"+Port, nil)
    
}
