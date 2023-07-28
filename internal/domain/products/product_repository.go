package products

import (
	"github.com/evermos/boilerplate-go/infras"
	"github.com/evermos/boilerplate-go/shared/logger"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	CreateWithVariant(payload ProductAndVariant) error
	GetAll()
}

type ProductRepositoryMariaDB struct {
	DB *infras.MariaDBConn
}

func ProvideProductRepositoryMariaDB(db *infras.MariaDBConn) *ProductRepositoryMariaDB {
	s := new(ProductRepositoryMariaDB)
	s.DB = db
	return s
}

func (r *ProductRepositoryMariaDB) CreateWithVariant(payload ProductAndVariant) error {
	return r.DB.WithTransaction(func(db *sqlx.Tx, c chan error) {
		if err := r.txCreateWithVariant(db, payload); err != nil {
			c <- err
			return
		}
		c <- nil
	})
}

func (r *ProductRepositoryMariaDB) txCreateWithVariant(tx *sqlx.Tx, payload ProductAndVariant) (err error) {

	query := `
		INSERT INTO product (product_id, product_name, brand_id, updated_at, created_by, created_at, updated_by,user_id)
		VALUES (:product_id, :product_name, :brand_id, :updated_at, :created_by,:created_at,:updated_by,:user_id);
	`
	varQuery := `INSERT INTO variant (variant_id,product_id,variant_name,price,quantity,updated_at, created_by, created_at, updated_by)
	 VALUES (:variant_id,:product_id,:variant_name,:price,:quantity,:updated_at, :created_by, :created_at, :updated_by)
	 `
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	varStmt, err := tx.PrepareNamed(varQuery)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	defer stmt.Close()
	defer varStmt.Close()
	_, err = stmt.Exec(payload.Product)
	if err != nil {
		logger.ErrorWithStack(err)
	}
	_, err = varStmt.Exec(payload.Variant)
	if err != nil {
		logger.ErrorWithStack(err)
	}

	return
}
