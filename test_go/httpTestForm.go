package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hey)
	http.ListenAndServe(":8080", nil)
}

const tpl = `
<html>
	<head>
		<title>Hey</title>
	</head>
	<body>
		<form method="post" action="/">
			Username:<input type="text" name="uname">
			Password:<input type="password" name="pwd">
			<button type="submit">Submit</button>
		</form>
	</body>
</html>
`

func Hey(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.New("hey")
		t.Parse(tpl)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println(r.FormValue("uname"))
	}
}
