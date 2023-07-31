package products

import (
	"fmt"

	"github.com/evermos/boilerplate-go/infras"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/shared/logger"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	CreateWithVariant(payload ProductAndVariant) error
	GetAllProducts(field, sort string, limit, offset int) (prods []Product, err error)
	GetProdcutByID(prodId string) (prod Product, err error)
	GetProductWithVariants(proId uuid.UUID) (prod ProductWithVariants, err error)
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

func (r *ProductRepositoryMariaDB) GetAllProducts(field, sort string, limit, offset int) (prods []Product, err error) {
	query := fmt.Sprintf("SELECT * FROM product ORDER BY %s %s LIMIT %d OFFSET %d", field, sort, limit, offset)
	err = r.DB.Read.Select(&prods, query)

	if err != nil {
		err = failure.InternalError(err)
		logger.ErrorWithStack(err)
		return
	}
	return
}

func (r *ProductRepositoryMariaDB) GetProdcutByID(prodId string) (prod Product, err error) {
	err = r.DB.Read.Get(&prod, "SELECT * FROM product WHERE product_id = %d", prodId)
	if err != nil {
		err = failure.InternalError(err)
		logger.ErrorWithStack(err)
		return
	}
	return
}

func (r *ProductRepositoryMariaDB) GetProductWithVariants(prodId uuid.UUID) (prod ProductWithVariants, err error) {
	err = r.DB.Read.Get(&prod.Product, "SELECT * FROM product WHERE product_id = ?", prodId.String())
	if err != nil {
		err = failure.InternalError(err)
		logger.ErrorWithStack(err)
		return
	}
	varQuery := fmt.Sprintf("SELECT * FROM variant WHERE product_id = '%s'", prodId.String())
	err = r.DB.Read.Select(&prod.Variants, varQuery)
	if err != nil {
		err = failure.InternalError(err)
		logger.ErrorWithStack(err)
		return
	}
	return
}
