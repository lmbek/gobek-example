package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var dataPath = "./../../data/"

type Database struct {
	connection *sql.DB
}

// Open a connection to the SQLite database.
func (database *Database) connect() error {
	var err error
	database.connection, err = sql.Open("sqlite3", dataPath+"gobek-example.db")
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) createTable(SQLFile string) error {
	query, err := os.ReadFile(SQLFile)
	if err != nil {
		return err
	}

	// Create the "user" table if it does not already exist.
	_, err = database.connection.Exec(string(query))
	if err != nil {
		return err
	}

	return nil
}

func (database *Database) insertData(SQLFile string, args ...any) error {
	query, err := os.ReadFile(SQLFile)
	if err != nil {
		return err
	}

	// Prepare the INSERT statement. (prepared statement)
	statement, err := database.connection.Prepare(string(query))
	defer statement.Close()
	if err != nil {
		return err
	}

	// Insert a new row into the "user" table using the prepared statement.
	result, err := statement.Exec(args...)
	if err != nil {
		fmt.Println("error inserting new user")
		return err
	}

	// Print the number of rows affected by the INSERT statement.
	fmt.Println(result.RowsAffected())

	return nil
}

// TODO: SELECT STATEMENT
func (database *Database) selectData(SQLFile string, args ...any) error {
	return errors.New("not implemented")
}

func (database *Database) updateData(SQLFile string, args ...any) error {
	query, err := os.ReadFile(SQLFile)
	if err != nil {
		return err
	}

	// Prepare the UPDATE statement. (prepared statement)
	statement, err := database.connection.Prepare(string(query))
	defer statement.Close()
	if err != nil {
		return err
	}

	// Modify the users data using the prepared statement.
	result, err := statement.Exec(args...)
	if err != nil {
		fmt.Println("error updating user")
		return err
	}

	// Print the number of rows affected by the UPDATE statement.
	fmt.Println(result.RowsAffected())

	return nil
}

func (database *Database) deleteData(SQLFile string, args ...any) error {
	query, err := os.ReadFile(SQLFile)
	if err != nil {
		return err
	}

	// Prepare the DELETE statement. (prepared statement)
	statement, err := database.connection.Prepare(string(query))
	defer statement.Close()
	if err != nil {
		return err
	}

	// Delete a new row into the "user" table using the prepared statement.
	result, err := statement.Exec(args...)
	if err != nil {
		return err
	}

	// Print the number of rows affected by the DELETE statement.
	fmt.Println(result.RowsAffected())

	return nil
}

func runExamples(database *Database) error {
	// Create users table
	err := database.createTable(dataPath + "sql/examples/create-table-users.sql")
	if err != nil {
		return err
	}

	// Insert data into users table
	err = database.insertData(dataPath+"sql/examples/insert-into-users.sql", nil, "lars2", 28, "")
	if err != nil {
		return err
	}

	// TODO add select data

	// UPDATE data from users table
	err = database.updateData(dataPath+"sql/examples/change-website-of-user.sql", "https://google.dk", 3)
	if err != nil {
		return err
	}

	// Remove data from users table
	err = database.deleteData(dataPath+"sql/examples/remove-from-users.sql", 4)
	if err != nil {
		return err
	}

	return nil
}

func run() error {
	// Initialize database
	database := new(Database)

	// Connect
	err := database.connect()
	defer database.connection.Close()
	if err != nil {
		return err
	}

	err = runExamples(database)
	if err != nil {
		return err
	}
	// We are done
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}
