package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/matus-kacmar/banka-web/sanitize"
)

var (
	logFile   *os.File
	router    *mux.Router
	templates map[string]*template.Template
)

type ErrorData struct {
	Message string
}

// Initialize all necessary things as log files and such...
func init() {
	var err error

	logFile, err = os.Open("logFile.txt")

	if !os.IsExist(err) {
		logFile, err = os.Create("logFile.txt")

		if err != nil {
			log.Fatal(err)
		}
	}

	makeLog("Server started at >>")

	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index"] = template.Must(template.ParseFiles("public/index.html"))
}

// Makes log to a log file
func makeLog(logInput string) {

	currentTime := time.Now().UTC().Format(time.RFC850)
	logMessage := logInput + " " + currentTime + "\n"

	logFile.WriteString(logMessage)

}

// Load html page from directory
func loadPage(pageName string, writer http.ResponseWriter, templateName string, viewModel interface{}) {
	templ, ok := templates[pageName]

	if !ok {
		http.Error(writer, "template does not exists", http.StatusInternalServerError)
	}

	if templateName == "" {
		err := templ.Execute(writer, viewModel)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err := templ.ExecuteTemplate(writer, templateName, viewModel)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Home page handler
func homePageHandler(writer http.ResponseWriter, request *http.Request) {
	errorMessage := request.URL.Query().Get("err")

	if errorMessage != "" {
		if strings.EqualFold(errorMessage, "login") {
			loadPage("index", writer, "", &ErrorData{Message: "Wrong username or password"})
		}
	} else {
		loadPage("index", writer, "", &ErrorData{Message: ""})
	}
}

// Login handler
func loginHandler(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("username")
	password := request.FormValue("password")

	if sanitize.ParseEmail(email) && sanitize.ParsePassword(password) {
		fmt.Fprint(writer, "Wohooo logged in!")
	} else {
		http.Redirect(writer, request, "/?err=login", 302)
	}
}

func main() {

	router = mux.NewRouter().StrictSlash(false)
	resourceFileServer := http.FileServer(http.Dir("./public/resources/"))
	router.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", resourceFileServer))

	router.HandleFunc("/", homePageHandler).Methods("GET")
	router.HandleFunc("/login", loginHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server listenning on port ", server.Addr)
	server.ListenAndServe()

}
