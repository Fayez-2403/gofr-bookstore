// pkg/config/app.go
package config

import (
	"github.com/Fayez-2403/gofr-bookstore/pkg/utils"
	"github.com/Fayez-2403/gofr-bookstore/pkg/models"
	"github.com/gofr-dev/gofr"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql`n`t" github.com/Fayez-2403/gofr-bookstore/pkg/models`n`t "github.com/Fayez-2403/gofr-bookstore/pkg/utils"
)

func ConfigureApp(app *gofr.App) {
	// Connect to the database
	Connect()

	// Register models with GoFr
	app.RegisterModel(&models.Book{})

	// Register utils with GoFr
	app.RegisterUtility(utils.ParseBody)
}

func Connect() {
	dsn := "Fayez:Fayez@tcp(localhost:8080)/simplerest?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	app.DB = d
}

func GetDB() *gorm.DB {
	return app.DB
}
