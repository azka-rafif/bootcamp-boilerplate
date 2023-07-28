package materials

import (
	"github.com/evermos/boilerplate-go/infras"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/shared/logger"
	"github.com/jmoiron/sqlx"
)

type MaterialRepository interface {
	Create(payload Material) error
	GetAll() (mats []Material, err error)
}

type MaterialRepositoryMySQL struct {
	DB *infras.MySQLConn
}

type MaterialRepositoryMariaDB struct {
	DB *infras.MariaDBConn
}

func ProvideMaterialRepositoryMySQL(db *infras.MySQLConn) *MaterialRepositoryMySQL {
	s := new(MaterialRepositoryMySQL)
	s.DB = db
	return s
}

func ProvideMaterialRepositoryMariaDB(db *infras.MariaDBConn) *MaterialRepositoryMariaDB {
	s := new(MaterialRepositoryMariaDB)
	s.DB = db
	return s
}

func (r *MaterialRepositoryMySQL) Create(payload Material) error {

	return r.DB.WithTransaction(func(db *sqlx.Tx, c chan error) {
		if err := r.txCreate(db, payload); err != nil {
			c <- err
			return
		}
		c <- nil
	})
}

func (r *MaterialRepositoryMySQL) txCreate(tx *sqlx.Tx, payload Material) (err error) {
	query := `
		INSERT INTO materials (id, title, description)
		VALUES (:id, :title, :description);
	`
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(payload)
	if err != nil {
		logger.ErrorWithStack(err)
	}

	return
}

func (r *MaterialRepositoryMariaDB) Create(payload Material) error {
	return r.DB.WithTransaction(func(db *sqlx.Tx, c chan error) {
		if err := r.txCreate(db, payload); err != nil {
			c <- err
			return
		}
		c <- nil
	})
}

func (r *MaterialRepositoryMariaDB) txCreate(tx *sqlx.Tx, payload Material) (err error) {
	query := `
		INSERT INTO materials (id, title, description)
		VALUES (:id, :title, :description);
	`
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(payload)
	if err != nil {
		logger.ErrorWithStack(err)
	}

	return
}

func (r *MaterialRepositoryMariaDB) GetAll() (mats []Material, err error) {
	err = r.DB.Read.Select(&mats, `select * from materials`)

	if err != nil {
		err = failure.InternalError(err)
		logger.ErrorWithStack(err)
		return
	}
	return

}
