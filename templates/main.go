package main

import (
	"net/http"
	"text/template"
)

type User struct {
	Username string
	Password string
}

var (
	// users       = make(map[string]string)

	users = map[string]string{
		"adil": "123123",
	}
	currentUser User
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	//Часть кода если мы пропускаем не зареганных(логин) пользователей
	// if currentUser.Username == "" {
	// 	http.Redirect(w, r, "./login", http.StatusSeeOther) // статус не будет остонавливать а дальше пропустит
	// 	return
	// }

	templ, err := template.ParseFiles("./index.html")
	if err != nil {
		http.Error(w, "Error loading index.html page", http.StatusInternalServerError)
		return
	}

	templ.Execute(w, map[string]interface{}{
		"User": currentUser,
	})
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templ, err := template.ParseFiles("./register.html")
		if err != nil {
			http.Error(w, "Error loading register.html page", http.StatusInternalServerError)
			return
		}

		templ.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			http.Error(w, "Username and Password are required", http.StatusBadRequest)
			return
		}

		if _, exist := users[username]; exist {
			http.Error(w, "User already exist", http.StatusBadRequest)
			return
		} else {
			users[username] = password
		}

		currentUser.Username = username
		currentUser.Password = password
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templ, err := template.ParseFiles("./login.html")
		if err != nil {
			http.Error(w, "Error loading login.html page", http.StatusInternalServerError)
			return
		}

		templ.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			http.Error(w, "Username and Password are required", http.StatusBadRequest)
			return
		}
		storedPassword, exist := users[username]
		if exist && storedPassword != password {
			http.Error(w, "Invalid password", http.StatusBadRequest)
			return
		} else if !exist {
			http.Error(w, "User is not exist", http.StatusBadRequest)
			return
		}

		currentUser.Username = username
		currentUser.Password = password

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	//TODO: logout
	http.HandleFunc("/", IndexPage)
	http.HandleFunc("/register", RegisterPage)
	http.HandleFunc("/login", LoginHandler)
	http.ListenAndServe(":8080", nil)
}
