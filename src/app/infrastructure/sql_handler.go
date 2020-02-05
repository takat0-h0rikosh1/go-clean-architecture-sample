package infrastructure

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
	"log"
	"my-echo.com/src/app/domain"
	"my-echo.com/src/app/interface/database"
)

type SQLHandler struct {
	Conn *gorp.DbMap
}

func newSQLHandler() database.SQLHandler {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	connStr := "postgres://my-postgres:my-postgres@localhost/my-postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(domain.Post{}, "Posts").SetKeys(false, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = dbmap
	return sqlHandler
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func (handler *SQLHandler) Execute(statement string, args ...interface{}) (err error) {
	_, err = handler.Conn.Exec(statement, args...)
	if err != nil {
		return
	}
	return
}

func (handler *SQLHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
