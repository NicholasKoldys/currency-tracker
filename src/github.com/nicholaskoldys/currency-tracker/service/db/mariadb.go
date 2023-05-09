package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nicholaskoldys/currency-tracker/util"
)

type DB struct{ *sql.DB }
type prepareStatement func(*sql.DB) (*sql.Stmt, error)

var dbInstance *sql.DB

/**
 */
func init() {

	util.Debug(9, "intializing Mariadb")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQLDB_USER"), os.Getenv("MYSQLDB_PASS"), os.Getenv("MYSQLDB_HOST"), os.Getenv("MYSQLDB_PORT"), os.Getenv("MYSQLDB_DB"))

	_, err := connectDB(connectionString)
	if err != nil {
		util.Debug(9, err.Error())
	}
}

/**
 */
func GetDB() *sql.DB {
	if isOpen := isDBConnected(); !isOpen {
		util.Debug(9, "Not Connected")
		return nil
	}
	return dbInstance
}

/**
 */
func connectDB(connectionString string) (*sql.DB, error) {

	util.Debug(9, "Opening MariaDB connection...")

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		util.Debug(9, err.Error())
		return nil, err
	}

	//! This must NOT be locally scoped. DoNot Use ":="
	dbInstance = db

	/* if err := dbInstance.Ping(); err != nil {
		util.Debug(9, err.Error())
		return nil, err
	} */

	util.Debug(9, "Success")

	/* * IF USING ORM WRAP HERE */

	return dbInstance, nil
}

/**
 */
func isDBConnected() bool {
	util.Debug(9, "Testing Connection")
	// Ping the database to check if the connection is successful
	if err := dbInstance.Ping(); err != nil {
		util.Debug(9, err.Error())
		return false
	}
	util.Debug(9, "Connected")
	return true
}

/**
 */
func TestCount() {
	var count int

	util.Debug(9, "Attemtping DB Count...")
	err := GetDB().QueryRow("SELECT COUNT(*) FROM currencyDecimals").Scan(&count)
	if err != nil {
		util.Debug(9, err.Error())
	}
}

/* import "golang.org/x/crypto/bcrypt"
import "os"

// GenerateHashedPassword generates a bcrypt hashed password from a plain-text password and a random salt.
func GenerateHashedPassword(password string) (string, error) {
    salt := make([]byte, 32)
    _, err := rand.Read(salt)
    if err != nil {
        return "", err
    }
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+string(salt)), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// CheckPasswordHash checks if a given plain-text password matches the given hashed password.
func CheckPasswordHash(password string, hashedPassword string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

// Store salt and hashed password in environment variables
os.Setenv("HASHED_PASSWORD", GenerateHashedPassword("my_password"));
os.Setenv("SALT", salt) */
