package product

import (
	"database/sql" //libreria para conectarse e interactuar con sql
	"errors"
	"time"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type repositoryDB struct {
	db *sql.DB // tiene un puntero sql
}

func NewRepositoryDB(db *sql.DB) Repository { //contructor recibe un puntero a una conexion sql
	return &repositoryDB{
		db: db,
	}
}

func (r *repositoryDB) GetByID(id int) (domain.Product, error) {
	query := "SELECT * FROM products WHERE id=?;"
	row := r.db.QueryRow(query, id)
	p := domain.Product{}
	err := row.Scan(&p.Id, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repositoryDB) Create(p domain.Product) (domain.Product, error) {

	fecha := p.Expiration
	t, _ := time.Parse("02/01/2006", fecha) //formatear fecha
	p.Expiration = t.Format("2006-01-02")
	query := "INSERT INTO products(name,quantity,code_value,is_published,expiration,price) VALUES (?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return domain.Product{}, err
	}

	res, err := stmt.Exec(p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price)
	if err != nil {
		return domain.Product{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}
	p.Id = int(id)
	return p, nil
}

func (r *repositoryDB) Update(id int, p domain.Product) (domain.Product, error) {
	query := "UPDATE products SET name=?, Quantity=?, Code_value=?, is_published=?, expiration=?, price=?  WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return domain.Product{}, err
	}

	res, err := stmt.Exec(p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.Id)
	if err != nil {
		return domain.Product{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repositoryDB) Delete(id int) error {
	query := "DELETE FROM products WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return errors.New("not found")
	}

	return nil
}
