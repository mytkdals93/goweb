package main

import (
	"html/template"
	"os"
)

/*
text/template 이냐 html/template이냐에 따라 특수문자 탈락이 결정됨
*/

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "상민", Email: "mytkdals93@naver.com", Age: 28}
	user2 := User{Name: "상민2", Email: "mytkdals93@gmail.com", Age: 40}
	users := []User{user, user2}
	// templ, err := template.New("Tem1").
	// 	Parse("Name: {{.Name}}\n" +
	// 		"Email: {{.Email}}\n" +
	// 		"Age: {{.Age}}\n")

	// if err != nil {
	// 	panic(err)
	// }
	// templ.Execute(os.Stdout, user)
	// templ.Execute(os.Stdout, user2)
	templ, err := template.New("Tem1").
		ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")

	if err != nil {
		panic(err)
	}
	// templ.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user)
	// templ.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user2)
	templ.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)

}
