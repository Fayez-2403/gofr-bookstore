// pkg/routes/bookstore-routes.go
package routes

import (
	"github.com/Fayez-2403/gofr-bookstore/pkg/controllers`n`t"github.com/Fayez-2403/gofr-bookstore/pkg/utils"
	"github.com/gofr-dev/gofr"
)

func RegisterBookStoreRoutes(r *gofr.Router) {
	r.POST("/book", controllers.CreateBook)
	r.GET("/book", controllers.GetBook)
	r.GET("/book/{bookId}", controllers.GetBookById)
	r.PUT("/book/{bookId}", controllers.UpdateBook)
	r.DELETE("/book/{bookId}", controllers.DeleteBook)
}
