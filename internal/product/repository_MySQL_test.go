package product

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func init() {
	dns := mysql.Config{
		User:   "root",
		Passwd: "",
		DBName: "my_db",
		Addr:   "127.0.0.1",
	}
	txdb.Register("txdb", "mysql", dns.FormatDSN())
}

func TestRepositoryMySQL_GetAll(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		//arrrange
		db, err := sql.Open("txdb", "identifier")
		assert.NoError(t, error)
		defer db.Close()
		rp := NewRepositoryDB(db)
		//act

		//exp:=[]*product{}

		//pr, err := rp.GetAll()

		//asssert
	})

}
