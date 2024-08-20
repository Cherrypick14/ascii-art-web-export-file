package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"artweb/asciiart"
)

type PageData struct {
	Result string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}
	// Render the form
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		NotFoundHandler(w, r)
		return
	}
	// Create an instance of PageData with the desired result
	data := PageData{
		Result: "", // Replace with your actual result(Ascii-art)
	}
	// Execute the template with the provided data
	if err := tpl.Execute(w, data); err != nil {
		NotFoundHandler(w, r)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		FourohFiveHandler(w, r)
		return
	}
	if r.URL.Path != "/ascii-art" {
		NotFoundHandler(w, r)
		return

	} else {
		// Process the form and generate ASCII art
		input := r.FormValue("message")
		input = strings.ReplaceAll(input, "\r", "\n")
		banner := r.FormValue("banner")
		if input == "" || banner == "" {
			BadRequestHandler(w, r)
			return
		}
		asciiArt, err := asciiart.Art(input, banner)
		if err != nil {
			if asciiArt == "400" {
				BadRequestHandler(w, r)

				return
			} else if asciiArt == "404" {
				NotFoundHandler(w, r)

				return
			}
		}
		data := PageData{
			Result: asciiArt,
		}
		w.WriteHeader(http.StatusOK)
		log.Println("200 OK: ASCII art generated successfully")
		_, eer := os.Stat("templates/index.html")
		if eer != nil {
			NotFoundHandler(w, r)

			return
		}
		// Render the template with the generated ASCII art
		tpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			InternalServerHandler(w, r)

			return

		}
		tpl.Execute(w, data)
	}
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/400.html")
	if err != nil {
		http.Error(w, "400 Bad Request ", http.StatusBadRequest)
		return
	}
	tpl.Execute(w, nil)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		http.Error(w, "404 Page Not Found ", http.StatusNotFound)
		return
	}
	tpl.Execute(w, nil)
}

func FourohFiveHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/405.html")
	if err != nil {
		http.Error(w, "405 Bad Method ", http.StatusMethodNotAllowed)
		return
	}
	tpl.Execute(w, nil)
}

func InternalServerHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/500.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error ", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}
