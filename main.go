package main

import (
	"encoding/gob"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/matus-kacmar/banka-web/database"
	"github.com/matus-kacmar/banka-web/sanitize"
	"golang.org/x/crypto/bcrypt"
)

var (
	logFile      *os.File
	router       *mux.Router
	sessionStore *sessions.CookieStore
	templates    map[string]*template.Template
	attempts     int8
)

// ErrorData is structure for parsing error messages to template.
type ErrorData struct {
	Message string
}

// Initialize all necessary things as log files and such...
func init() {
	var err error
	attempts = 0

	logFile, err = os.Open("logFile.txt")

	if !os.IsExist(err) {
		logFile, err = os.Create("logFile.txt")

		if err != nil {
			log.Fatal(err)
		}
	}

	sessionStore = sessions.NewCookieStore([]byte("much-secret-phrase-very-secure"))
	gob.Register(&database.Client{})
	makeLog("Server started at >>")

	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index"] = template.Must(template.ParseFiles("public/index.html"))
	templates["bank"] = template.Must(template.ParseFiles("public/pages/bank.html", "public/templates/header.html", "public/templates/main.html", "public/templates/footer.html"))
	templates["accounts"] = template.Must(template.ParseFiles("public/pages/accounts.html"))
	templates["loans"] = template.Must(template.ParseFiles("public/pages/loans.html"))
	templates["transactions"] = template.Must(template.ParseFiles("public/pages/transactions.html"))
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
	// try to get client data from session if they exist
	session, _ := sessionStore.Get(request, "bank-user")
	data := session.Values["user"]
	_, ok := data.(*database.Client)

	// If client data in session was found redirect to bank page
	if ok {
		http.Redirect(writer, request, "/bank", 302)
	}

	errorMessage := request.URL.Query().Get("err")

	if errorMessage != "" {

		if strings.EqualFold(errorMessage, "login") {
			loadPage("index", writer, "", &ErrorData{Message: "Wrong username or password"})
		} else if strings.EqualFold(errorMessage, "wrong-settings") {
			loadPage("index", writer, "", &ErrorData{Message: "Wrong browser settings cannot log in"})
		} else if strings.EqualFold(errorMessage, "blocked") {
			loadPage("index", writer, "", &ErrorData{Message: "Too many wrong attempts your account is blocked"})
		}

	} else {
		loadPage("index", writer, "", &ErrorData{Message: ""})
	}
}

// Login handler
func loginHandler(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("username")
	password := request.FormValue("password")

	// Compare input to sanitize package regexes
	if sanitize.ParseEmail(email) && sanitize.ParsePassword(password) {
		client := database.GetClientByUsername(email)
		err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password))

		// err == nil means if password entered by user matches password from db
		if err == nil && client.Active {
			// store session with user info
			session, _ := sessionStore.Get(request, "bank-user")
			session.Values["user"] = &client

			// sets session otions, age o session is set to one minute for testing purposes
			session.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   3600,
				HttpOnly: true,
			}

			err := session.Save(request, writer)

			if err != nil {
				log.Println(err)
				http.Redirect(writer, request, "/?err=wrong-settings", 302)
			}

			http.Redirect(writer, request, "/bank", 302)
		} else {
			// If account is innactive display user block message
			if !client.Active {
				http.Redirect(writer, request, "/?err=blocked", 302)
			}
			// If wrong attempt add up wrong attempts
			if attempts < 3 {
				attempts++
			} else {
				// If more than 3 attempts block user and display message
				if database.BlockUser(client.ID) {
					http.Redirect(writer, request, "/?err=blocked", 302)
				}
			}

			// Else display wrong username or password message
			http.Redirect(writer, request, "/?err=login", 302)
		}
	} else {
		http.Redirect(writer, request, "/?err=login", 302)
	}
}

func bankHandler(writer http.ResponseWriter, request *http.Request) {
	session, _ := sessionStore.Get(request, "bank-user")
	data := session.Values["user"]
	client, ok := data.(*database.Client)

	if !ok {
		http.Redirect(writer, request, "/", 302)
	}

	loadPage("bank", writer, "", client)
}

// logOutHandler destroys client session and redirects client to login page
func logOutHandler(writer http.ResponseWriter, request *http.Request) {
	session, _ := sessionStore.Get(request, "bank-user")
	session.Values["user"] = nil
	err := session.Save(request, writer)

	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/?err=wrong-settings", 302)
	}

	http.Redirect(writer, request, "/", 302)
}

// accountHandler response with account data
func accountHandler(writer http.ResponseWriter, request *http.Request) {
	session, _ := sessionStore.Get(request, "bank-user")
	data := session.Values["user"]
	client, ok := data.(*database.Client)

	if !ok {
		http.Redirect(writer, request, "/", 302)
	}

	accounts := database.GetClientAccountsById(client.ID)
	loadPage("accounts", writer, "", accounts)
}

// accountHandler response with account data
func loansHandler(writer http.ResponseWriter, request *http.Request) {
	session, _ := sessionStore.Get(request, "bank-user")
	data := session.Values["user"]
	client, ok := data.(*database.Client)

	if !ok {
		http.Redirect(writer, request, "/", 302)
	}

	loans := database.GetClientLoansById(client.ID)
	loadPage("loans", writer, "", loans)
}

// accountHandler response with account data
func transactionsHandler(writer http.ResponseWriter, request *http.Request) {
	session, _ := sessionStore.Get(request, "bank-user")
	data := session.Values["user"]
	client, ok := data.(*database.Client)

	if !ok {
		http.Redirect(writer, request, "/", 302)
	}

	transactions := database.GetClientTransactionsById(client.ID)
	loadPage("transactions", writer, "", transactions)
}

func main() {

	router = mux.NewRouter().StrictSlash(false)
	resourceFileServer := http.FileServer(http.Dir("./public/resources/"))
	router.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", resourceFileServer))

	router.HandleFunc("/", homePageHandler).Methods("GET")
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/bank", bankHandler).Methods("GET")
	router.HandleFunc("/bank/logout", logOutHandler).Methods("GET")
	router.HandleFunc("/bank/accounts", accountHandler).Methods("GET")
	router.HandleFunc("/bank/loans", loansHandler).Methods("GET")
	router.HandleFunc("/bank/transactions", transactionsHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server listenning on port ", server.Addr)
	server.ListenAndServe()

}
