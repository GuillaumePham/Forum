package main
import(
	"net/http"
)
/*func compte(w http.ResponseWriter, r *http.Request){
	email := r.FormValue("mail")
	pseudo := r.FormValue("username")
	password := r.FormValue("password")
	Checkpassword := r.FormValue("password2")
	if(password == Checkpassword){
		if (len(password) >= 8){
			requestDB(pseudo, email, password)
		}
	}
    tmpl.ExecuteTemplate(w, "account", nil)
}
func requestDB(pseudo, email, password string){
	rows, _ := db.Query("INSERT INTO `users` (`id`, `pseudo`, `adresse_mail`, `motdepasse`) VALUES (DEFAULT, ?, ?, ?)",pseudo, email, password)
	defer rows.Close()
}
*/