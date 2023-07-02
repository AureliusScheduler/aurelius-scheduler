package db_context

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type DbOptions struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type DbContext struct {
	Db *sql.DB
}

func NewDbContext(options *DbOptions) DbContext {
	db, err := sql.Open("mysql", options.User+":"+options.Password+"@tcp("+options.Host+":"+options.Port+")/"+options.Database)

	if err != nil {
		panic(err.Error())
	}

	return DbContext{
		Db: db,
	}
}
