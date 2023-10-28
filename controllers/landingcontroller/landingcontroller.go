package landingcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Selamat Datang di Golang RestAPI", "Create by": "Muhammad Iqbal", "Host": "/api/books"})
}
