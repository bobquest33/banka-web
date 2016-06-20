package database

import (
	"database/sql"
	"log"

	// Import of postgresql driver
	_ "github.com/lib/pq"
)

// Data for database driver
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

// Account struct holds cllients account data
type Account struct {
	AccountID int
	Balance   float64
}

// Loan struct holds clients loan data
type Loan struct {
	Amount     float64
	PaidAmount float64
	Interest   float64
}

// Transaction struct holds the data of client transaction
type Transaction struct {
	ClientRequest bool
	AccountID     int
	TransDate     string
	Value         float64
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

// GetClientAccountsByID returns array of account structs with client data
func GetClientAccountsByID(id int) []Account {
	query := "SELECT accountid,balance FROM accounts WHERE clientid=$1"

	statement, err := database.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	result, err := statement.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var (
		accounts  []Account
		accountid int
		balance   float64
	)

	for result.Next() {
		result.Scan(&accountid, &balance)
		accounts = append(accounts, Account{accountid, balance})
	}

	return accounts
}

// GetClientLoansByID return list of loans of the client
func GetClientLoansByID(id int) []Loan {
	query := "SELECT amount,paidamount,interest FROM loans WHERE clientid=$1"

	statement, err := database.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	result, err := statement.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var (
		loans      []Loan
		amount     float64
		paidamount float64
		interest   float64
	)

	for result.Next() {
		result.Scan(&amount, &paidamount, &interest)
		loans = append(loans, Loan{amount, paidamount, interest})
	}

	return loans
}

// GetClientTransactionsByID return list of users transactions
func GetClientTransactionsByID(id int) []Transaction {
	query := "SELECT clientrequest,accountid,transdate,value FROM transactions WHERE personid=$1"

	statement, err := database.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()

	result, err := statement.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	var (
		transactions  []Transaction
		clientrequest bool
		accountid     int
		transdate     string
		value         float64
	)

	for result.Next() {
		result.Scan(&clientrequest, &accountid, &transdate, &value)
		log.Println(clientrequest, " ", accountid, " ", transdate, " ", value)
		transactions = append(transactions, Transaction{clientrequest, accountid, transdate, value})
	}

	return transactions
}

// BlockUser blocks user
func BlockUser(id int) bool {
	query := "UPDATE clientlogin SET active=$1 WHERE id=$2"

	statement, err := database.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	result, err := statement.Exec(false, id)

	if err != nil {
		return false
	}

	affected, err := result.RowsAffected()

	if err != nil && affected == 0 {
		return false
	}

	return true
}

// CloseConnection terminates connection to database
func CloseConnection() {
	database.Close()
}
