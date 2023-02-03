package main

import (
	"fmt"
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
		pageSize := c.DefaultQuery("max", "6")
		var covers []Cover
		query := fmt.Sprintf("SELECT * FROM covers WHERE deleted_at IS NULL ORDER BY 1 LIMIT %s", pageSize)
		db.Raw(query).Scan(&covers)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"covers": covers,
		})
	})
	r.Run()
}
