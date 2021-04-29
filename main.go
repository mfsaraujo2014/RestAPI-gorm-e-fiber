package main

import (
	"fmt"
	"modulo/book"
	"modulo/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Olá")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Falha ao conectar ao banco")
	}
	fmt.Println("O banco de dados foi aberto com sucesso")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("O banco de dados foi migrado")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}
