package materials

import (
	"database/sql"

	"github.com/evermos/boilerplate-go/infras"
	"github.com/evermos/boilerplate-go/shared/logger"
	"github.com/jmoiron/sqlx"
)

type MaterialRepository interface {
	Create(payload PayloadMaterial) error
	GetAll(field, sort string, limit, offset int) (*sql.Rows, error)
}

type MaterialRepositoryImpl struct {
	DB *infras.MySQLConn
}

func ProvideMaterialRepositoryMySQL(db *infras.MySQLConn) *MaterialRepositoryImpl {
	s := new(MaterialRepositoryImpl)
	s.DB = db
	return s
}

func (r *MaterialRepositoryImpl) Create(payload PayloadMaterial) error {

	return r.DB.WithTransaction(func(db *sqlx.Tx, c chan error) {
		if err := r.txCreate(db, payload); err != nil {
			c <- err
			return
		}
		c <- nil
	})
}

func (r *MaterialRepositoryImpl) txCreate(tx *sqlx.Tx, payload PayloadMaterial) (err error) {
	query := `
		INSERT INTO materials (title, description)
		VALUES (:title, :description);
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
