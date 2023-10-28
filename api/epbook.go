package api

import (
	"net/http"

	"github.com/Miqbal20/golang-restapi/controllers/bookcontroller"
	"github.com/Miqbal20/golang-restapi/models"
	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func Route(r *gin.RouterGroup) {
	models.ConnectDB()

	r.GET("/index", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/book", bookcontroller.Index)
	r.GET("/book/:id", bookcontroller.Show)
	r.POST("/book", bookcontroller.Create)
	r.PATCH("/book/:id", bookcontroller.Update)
	r.DELETE("/book", bookcontroller.Delete)
}

// func init() {
// 	app = gin.New()
// 	r := app.Group("/api")
// 	Route(r)
// }

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
