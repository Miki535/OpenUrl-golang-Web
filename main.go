package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", backendfunc)
	http.ListenAndServe(":8080", nil)
}

func Open(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func backendfunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		URL := r.FormValue("URL")
		Open(URL)
	}
	tpl.Execute(w, nil)
}

//test
