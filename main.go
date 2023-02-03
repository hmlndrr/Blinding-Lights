package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Cover struct {
	gorm.Model
	Url   string `json:"url"`
	Label string `json:"label"`
	gorm.DeletedAt
}

func main() {

	db, err := gorm.Open(sqlite.Open("covers.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Cover{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		q := c.DefaultQuery("q", "1")
		var covers []Cover
		db.Order(q).Find(&covers)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"covers": covers,
		})
	})
	r.Run()
}
