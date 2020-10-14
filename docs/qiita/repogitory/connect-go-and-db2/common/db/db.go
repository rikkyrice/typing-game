package db

import (
	"database/sql"
	"fmt"

	// Get a connection to db2 container
	_ "github.com/ibmdb/go_ibm_db"
)

// DBConn definition of connection
type DBConn struct {
	DB *sql.DB
	Tx *sql.Tx
}

// DBConnConfig settings for connection
type DBConnConfig struct {
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// NewDBConn create DB connection
func NewDBConn(c *DBConnConfig) (*DBConn, error) {
	conn := "HOSTNAME=" + c.Hostname + ";DATABASE=" + c.Database + ";PORT=" + c.Port + ";UID=" + c.Username + ";PWD=" + c.Password
	db, err := sql.Open("go_ibm_db", conn)
	if err != nil {
		fmt.Println("failed to open.")
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("failed to ping.")
		return nil, err
	}
	return &DBConn{DB: db}, nil
}

// Close close db connection
func (conn *DBConn) Close() {
	conn.DB.Close()
}

// BeginTx start transaction
func (conn *DBConn) BeginTx() error {
	tx, err := conn.DB.Begin()
	if err != nil {
		return err
	}
	conn.Tx = tx
	return nil
}

// Rollback rollback
func (conn *DBConn) Rollback() {
	if conn.Tx != nil {
		conn.Tx.Rollback()
	}
	conn.Tx = nil
}

// Commit commit
func (conn *DBConn) Commit() {
	if conn.Tx != nil {
		conn.Tx.Commit()
	}
	conn.Tx = nil
}

// GetPstmt getting statement
func (conn *DBConn) GetPstmt(stmt string) (*sql.Stmt, error) {
	if conn.Tx == nil {
		return conn.DB.Prepare(stmt)
	}
	return conn.Tx.Prepare(stmt)
}

// ConvertPstmt contain statement into transaction if Tx exists
func (conn *DBConn) ConvertPstmt(pstmt *sql.Stmt) *sql.Stmt {
	if conn.Tx == nil {
		return pstmt
	}
	return conn.Tx.Stmt(pstmt)
}
