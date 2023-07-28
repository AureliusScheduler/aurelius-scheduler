package db_context

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

type DbOptions struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type DbContext struct {
	Db     *sql.DB
	GormDb *gorm.DB
}

func NewDbContext(options *DbOptions) *DbContext {
	db, err := sql.Open("mysql", options.User+":"+options.Password+"@tcp("+options.Host+":"+options.Port+")/"+options.Database+"?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "aur_",
			SingularTable: true,
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             0,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
				Colorful:                  true,
			},
		),
	})

	if err != nil {
		panic(err.Error())
	}

	return &DbContext{
		Db:     db,
		GormDb: gormDb,
	}
}
