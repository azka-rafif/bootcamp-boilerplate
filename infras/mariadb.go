package infras

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/evermos/boilerplate-go/shared/failure"

	"github.com/evermos/boilerplate-go/configs"
	// use MariaDB driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

// MariaDBConn wraps a pair of read/write MariaDB connections.
type MariaDBConn struct {
	Read  *sqlx.DB
	Write *sqlx.DB
}

// ProvideMariaDBConn is the provider for MariaDBConn.
func ProvideMariaDBConn(config *configs.Config) *MariaDBConn {
	return &MariaDBConn{
		Read:  CreateMariaDBReadConn(*config),
		Write: CreateMariaDBWriteConn(*config),
	}
}

// CreateMariaDBWriteConn creates a database connection for write access.
func CreateMariaDBWriteConn(config configs.Config) *sqlx.DB {
	return CreateDBConnection(
		"write",
		config.DB.MariaDB.Write.Username,
		config.DB.MariaDB.Write.Password,
		config.DB.MariaDB.Write.Host,
		config.DB.MariaDB.Write.Port,
		config.DB.MariaDB.Write.Name,
		config.DB.MariaDB.Write.Timezone)

}

// CreateMariaDBReadConn creates a database connection for read access.
func CreateMariaDBReadConn(config configs.Config) *sqlx.DB {
	return CreateDBConnection(
		"read",
		config.DB.MariaDB.Read.Username,
		config.DB.MariaDB.Read.Password,
		config.DB.MariaDB.Read.Host,
		config.DB.MariaDB.Read.Port,
		config.DB.MariaDB.Read.Name,
		config.DB.MariaDB.Read.Timezone)

}

// CreateDBConnection creates a database connection.
func CreateMariaDBConnection(name, username, password, host, port, dbName, timeZone string) *sqlx.DB {
	descriptor := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s&parseTime=true",
		username,
		password,
		host,
		port,
		dbName,
		url.QueryEscape(timeZone))
	db, err := sqlx.Connect("mysql", descriptor)
	if err != nil {
		log.
			Fatal().
			Err(err).
			Str("name", name).
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Failed connecting to database")
	} else {
		log.
			Info().
			Str("name", name).
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Connected to database")
	}
	db.SetMaxIdleConns(maxIdleConnection)
	db.SetMaxOpenConns(maxOpenConnection)

	return db
}

// OpenMock opens a database connection for mocking purposes.
func OpenMockMariaDB(db *sql.DB) *MariaDBConn {
	conn := sqlx.NewDb(db, "mysql")
	return &MariaDBConn{
		Write: conn,
		Read:  conn,
	}
}

// WithTransaction performs queries with transaction
func (m *MariaDBConn) WithTransaction(block Block) (err error) {
	e := make(chan error)
	tx, err := m.Write.Beginx()
	if err != nil {
		return
	}
	go block(tx, e)
	err = <-e
	if err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			err = failure.InternalError(errTx)
		}
		return
	}
	err = tx.Commit()
	return
}
