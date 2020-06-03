package controller

import (
	"github.com/gofiber/fiber"
)


//GetBooks to Get all books
func GetProducts(c *fiber.Ctx)  {
	
	c.JSON("Got")
}

// //GetBook to Get one book
// func GetBook(c *fiber.Ctx)  {
// 	id:= c.Params("id")
// 	db := database.DBConn
// 	var book Book
// 	db.Find(&book,id)
// 	c.JSON(book)
// }

// //NewBook to Create a new book
// func NewBook(c *fiber.Ctx)  {
// 	db := database.DBConn
// 	book := new(Book)
// 	if err := c.BodyParser(book); err != nil {
// 		c.Status(503).Send(err)
// 		return
// 	}
//  	db.Create(&book)
// 	c.JSON(book)
// }

// //DeleteBook to Delete a book
// func DeleteBook(c *fiber.Ctx)  {
// 	id:= c.Params("id")
// 	db := database.DBConn
// 	var book Book
// 	db.First(&book,id)
// 	if book.Title == ""{
// 		c.Status(500).Send("No book found with given ID")
// 		return
// 	}
// 	db.Delete(&book)
// 	c.Send("Book successfully deleted")
// }