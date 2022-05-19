package main
import(
	"net/http"
	"log"
)
var tmpl = template.Must(template.ParseFiles("static/html/compte.html"))
func compte(w http.ResponseWriter, r *http.Request){
	email := r.FormValue("mail")
	pseudo := r.FormValue("username")
	password := r.FormValue("password")
	Checkpassword := r.FormValue("password2")
	if(password == Checkpassword){
		if (len(password) >= 8){
			requestDB()
		}
	}
	tmpl.Execute(w, tmpl)
}
func requestDB(pseudo,email,password){
	rows, _:= db.Query("INSERT INTO `Users` (`id`, `Pseudo`, `adresse_mail`, `motdepasse`) VALUES (DEFAULT,"+ pseudo +","+ email +","+ 4563 ")")
	defer rows.Close()
}
