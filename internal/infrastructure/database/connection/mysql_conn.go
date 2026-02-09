package connection

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLDB(username, password, host, port, dbName string) (*sql.DB, error) {

	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = username
	mysqlConfig.Passwd = password
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = fmt.Sprintf("%s:%s", host, port)
	mysqlConfig.DBName = dbName
	// Ensure DATETIME/TIMESTAMP columns are returned as time.Time values
	// so scanning into *time.Time works (e.g. task.deadline, created_at, updated_at).
	mysqlConfig.ParseTime = true

	// Optionally skip TLS verification when connecting to the database.
	// If the environment variable `DB_TLS_SKIP_VERIFY` is set to "true" (case-sensitive),
	// register a TLS config that disables certificate verification and instruct the
	// MySQL driver to use it.
	if os.Getenv("DB_TLS_SKIP_VERIFY") == "true" {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}
		// Register a named TLS config with the MySQL driver.
		if err := mysql.RegisterTLSConfig("skip-verify", tlsConfig); err == nil {
			mysqlConfig.TLSConfig = "skip-verify"
		}
	}

	connString := mysqlConfig.FormatDSN()
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, fmt.Errorf("open mysql: %w", err)
	}

	// Pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

func CheckConnectivity(db *sql.DB) error {

	pingTest := db.Ping()
	if pingTest != nil {
		return fmt.Errorf("Connection Failed: %w", pingTest)
	}
	return nil
}
