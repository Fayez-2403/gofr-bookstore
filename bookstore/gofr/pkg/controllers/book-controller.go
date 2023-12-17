// pkg/controllers/book-controller.go
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Fayez-2403/gofr-bookstore/pkg/models"
	"github.com/Fayez-2403/gofr-bookstore/pkg/utils"
	"github.com/gofr-dev/gofr"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(ctx *gofr.Context) (interface{}, error) {
	newBooks := models.GetAllBooks()
	return newBooks, nil
}

func GetBookById(ctx *gofr.Context) (interface{}, error) {
	bookId, err := strconv.ParseInt(ctx.PathParam("bookId"), 0, 0)
	if err != nil {
		return nil, err
	}

	bookDetails, err := models.GetBookById(bookId)
	if err != nil {
		return nil, err
	}

	return bookDetails, nil
}

func CreateBook(ctx *gofr.Context) (interface{}, error) {
	createBook := &models.Book{}
	utils.ParseBody(ctx.Request(), createBook)
	b := createBook.CreateBook()
	return b, nil
}

func DeleteBook(ctx *gofr.Context) (interface{}, error) {
	bookId, err := strconv.ParseInt(ctx.PathParam("bookId"), 0, 0)
	if err != nil {
		return nil, err
	}

	book := models.DeleteBook(bookId)
	return book, nil
}

func UpdateBook(ctx *gofr.Context) (interface{}, error) {
	updateBook := &models.Book{}
	utils.ParseBody(ctx.Request(), updateBook)
	bookId, err := strconv.ParseInt(ctx.PathParam("bookId"), 10, 64)
	if err != nil {
		return nil, err
	}

	bookDetails, _ := models.GetBookById(bookId)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	models.GetDB().Save(&bookDetails)

	return bookDetails, nil
}
