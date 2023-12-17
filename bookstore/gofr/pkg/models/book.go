// pkg/models/book.go
package models

import (
	"github.com/Fayez-2403/gofr-bookstore/pkg/config"
	"github.com/gofr-dev/gofr"
	"github.com/jinzhu/gorm"
)

var app *gofr.App
var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	app = gofr.GetApp()
	config.ConfigureApp(app)
	db = app.DB
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, error) {
	var getBook Book
	err := db.Where("ID=?", Id).Find(&getBook).Error
	return &getBook, err
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
