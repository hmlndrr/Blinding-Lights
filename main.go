package main

import (
	"net/http"
	"strings"
	"text/template"

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

	var covers []Cover
	db.Unscoped().Find(&covers)

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"covers": covers,
		})
	})
	r.Run()
}
