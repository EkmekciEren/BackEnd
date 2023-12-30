package main

import (
	"fmt"
	"net/http"
)

func IsEmpty(data string) bool {
	return len(data) == 0
}

func main() {
	mux := http.NewServeMux()

	// Kullanıcıdan alacağımız bilgiler
	uName, email, pwd, pwdConfirm := "", "", "", ""

	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hoşgeldiniz, lütfen kayıt olun veya giriş yapın.")
	})
	//signup
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")

		uNameCheck := IsEmpty(uName)
		emailCheck := IsEmpty(email)
		pwdCheck := IsEmpty(pwd)
		pwdConfirmCheck := IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "ErrorCode is -10: There is empty data.")
			return
		}
		if pwd == pwdConfirm {
			fmt.Fprintf(w, "Registration successful")
		} else {
			fmt.Fprintf(w, "Password information must be the same.")
		}
	})

	//log in
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		pwd = r.FormValue("password")
		r.ParseForm()

		email = r.FormValue("email")
		pwdConfirm := r.FormValue("confirm")

		emailCheck := IsEmpty(email)
		pwdCheck := IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintln(w, "Email and/or Password can not be empty")
		}
		if pwd != pwdConfirm {
			fmt.Println("Passwords do not match!")
		}
		dbEmail := "example@example.com"
		dbPwd := "12345!"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintln(w, "Login successfuly.")
		} else {
			fmt.Fprintln(w, "Login failed.")
		}
	})

	http.ListenAndServe(":8080", mux)
}
