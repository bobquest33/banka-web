package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	logFile *os.File
)

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
}

// Makes log to a log file
func makeLog(logInput string) {

	currentTime := time.Now().UTC().Format(time.RFC850)
	logMessage := logInput + " " + currentTime + "\n"

	logFile.WriteString(logMessage)

}

// Load html page from directory
func loadPage(pageName string, writer http.ResponseWriter) []byte {
	file, err := ioutil.ReadFile("public/" + pageName + ".html")

	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	return file
}

// Home page handler
func homePageHandler(writer http.ResponseWriter, request *http.Request) {
	page := loadPage("index", writer)

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(page))
}

func loginHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	username := vars["username"]
	password := vars["password"]

	log.Println(username + " " + password)
}

func main() {

	router := mux.NewRouter()
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
