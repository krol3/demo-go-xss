package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
)

const queryKey = "myKey"

type User struct {
	ID       int
	Email    string
	Password string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Demo home page")
}

func xssInsecureBasic(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get(queryKey))
	fmt.Fprintf(w, "Demo xssInsecureBasic page")

}

func xssInsecure(w http.ResponseWriter, r *http.Request) {

	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}

	value := query.Get(queryKey)
	if len(value) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing parameter ?%s=xxxxx", queryKey)
		return
	}

	var user1 = &User{1, "my.email@gmail.com", "golang.brazil#password"}
	var tmpl = fmt.Sprintf(`
<html>
<head>
<title>my-html</title>
</head>
<h1>Golang Brazil 2022</h1>
<h2>No search results for %s</h2>
<h2> Hi {{ .Email }}<h2>
</html>`, r.URL.Query().Get(queryKey))

	t, err := template.New("page").Parse(tmpl)

	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, &user1)
	fmt.Fprintf(w, "\n Demo XSS ParseQuery 01  - Insecure \n")
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/xss/insecure01", xssInsecureBasic)
	http.HandleFunc("/xss/insecure02", xssInsecure)

	http.ListenAndServe(":8085", nil)
}
