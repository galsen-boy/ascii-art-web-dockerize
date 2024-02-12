package main

import (
	"net/http"
	"os"
	"text/template"
)

// Starting to find the html page
func StartPage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" || req.Method != "GET" {
		error404(res)
		return
	} else {
		pageTemplate, err := template.ParseFiles("./templates/home.page.tmpl")
		if err != nil {
			error500(res)
			return
		}
		res.WriteHeader(200)
		pageTemplate.Execute(res, nil)
	}
}

// Form data after submiting

func SubmitTing(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		error405(res)
		return
	}
	if req.URL.Path != "/ascii-art" {
		error404(res)
		return
	}
	tpl, err := template.ParseFiles("./templates/home.page.tmpl")
	req.ParseForm()
	if req.FormValue("banner") != "standard" &&
		req.FormValue("banner") != "thinkertoy" &&
		req.FormValue("banner") != "shadow" || err != nil {
		error500(res)
		return
	}

	txt := req.FormValue("input")
	ban := req.FormValue("banner")
	i, _ := os.ReadFile(ban + ".txt")
	if ban == "shadow" {
		if len(i) != 7463 {
			error500(res)
			return
		}
	}
	if ban == "standard" {
		if len(i) != 26488 {
			error500(res)
			return
		}
	}
	if ban == "thinkertoy" {
		if len(i) != 4703 {
			error500(res)
			return
		}
	}

	_, err = os.ReadFile(ban + ".txt")
	if err != nil {
		error500(res)
		return
	}
	// isValid find if there an anpritable caracter
	if !isValid(txt) {
		error400(res)
		return
	}
	if err != nil {
		// log.Fatalln(err)
		error500(res)
		return
	}
	res.WriteHeader(200)
	tpl.Execute(res, ConvertStr(txt, ban))

}
