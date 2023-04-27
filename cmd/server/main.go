package main

import (
	"database/sql"
	"fmt"

	"github.com/bootcamp-go/consignas-go-db.git/cmd/server/handler"
	"github.com/bootcamp-go/consignas-go-db.git/internal/product"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	//"github.com/bootcamp-go/consignas-go-db.git/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	godotenv.Load()
	db, err := sql.Open("mysql", "root:@/my_db")

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Print(err)
	}
	//storage := store.NewJsonStore("./products.json")

	repo := product.NewRepositoryDB(db)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products/")
	{
		products.GET(":id", productHandler.GetByID())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
		products.GET("reportProducts/:id", productHandler.GetByID())
	}

	warehouses := r.Group("/warehouses/")
	{
		warehouses.GET(":id", productHandler.GetByID())
		warehouses.POST("", productHandler.Post())
	}

	r.Run(":8080")
}
