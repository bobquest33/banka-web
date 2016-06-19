package database

import (
	"database/sql"
	"log"

	// Import of postgresql driver
	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "sovy-bank"
	sslMode    = "disable"
)

var (
	database *sql.DB
	query    string
)

// Client structs hold client data from the database
type Client struct {
	ID          int
	FirstName   string
	LastName    string
	DateOfBirth string
	Username    string
	Password    string
	Active      bool
}

func init() {
	var err error

	connectionString := "user=" + dbUser + " dbname=" + dbName + " sslmode=" + sslMode
	database, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}
}

// GetClientByUsername returns a client struct with client data
func GetClientByUsername(clientUsername string) Client {
	query = "SELECT clients.id,clients.firstname,clients.lastname,clients.dateofbirth,clientlogin.username,clientlogin.password,clientlogin.active FROM clientlogin INNER JOIN clients ON clientlogin.id=clients.id WHERE clientlogin.username=$1"

	statement, err := database.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	var (
		id          int
		firstname   string
		lastname    string
		dateofbirth string
		username    string
		password    string
		active      bool
	)

	err = statement.QueryRow(clientUsername).Scan(&id, &firstname, &lastname, &dateofbirth, &username, &password, &active)

	if err != nil {
		log.Fatal(err)
	}

	return Client{id, firstname, lastname, dateofbirth, username, password, active}
}

// CloseConnection terminates connection to database
func CloseConnection() {
	database.Close()
}
