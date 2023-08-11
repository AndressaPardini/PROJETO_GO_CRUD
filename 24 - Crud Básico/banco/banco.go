package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

// Database é uma interface que representa uma conexão com o banco de dados
type Database interface {
	Prepare(query string) (Statement, error)
	Close() error
}

// Statement é uma interface que representa uma instrução SQL preparada
type Statement interface {
	Exec(args ...interface{}) (Result, error)
	Close() error
}

// Result é uma interface que representa o resultado de uma operação SQL
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// DB implementa a interface Database usando *sql.DB
type DB struct {
	*sql.DB
}

// Stmt implementa a interface Statement usando *sql.Stmt
type Stmt struct {
	*sql.Stmt
}

// Res implementa a interface Result usando *sql.Result
type Res struct {
	sql.Result
}

// Conectar abre a conexão com o banco de dados
func Conectar() (Database, error) {
	stringConexao := "Projetogo:Golang1!@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return &DB{db}, nil
}

// Prepare implementa o método Prepare da interface Database
func (db *DB) Prepare(query string) (Statement, error) {
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	return &Stmt{stmt}, nil
}

// Exec implementa o método Exec da interface Statement
func (stmt *Stmt) Exec(args ...interface{}) (Result, error) {
	result, err := stmt.Stmt.Exec(args...)
	if err != nil {
		return nil, err
	}
	return &Res{result}, nil
}

// Close implementa o método Close da interface Statement
func (stmt *Stmt) Close() error {
	return stmt.Stmt.Close()
}

// LastInsertId implementa o método LastInsertId da interface Result
func (res *Res) LastInsertId() (int64, error) {
	return res.Result.LastInsertId()
}

// RowsAffected implementa o método RowsAffected da interface Result
func (res *Res) RowsAffected() (int64, error) {
	return res.Result.RowsAffected()
}
