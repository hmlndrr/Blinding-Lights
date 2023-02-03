package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func Init() {
	db, err := gorm.Open(sqlite.Open("covers.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Cover{})

	db.Model(&Cover{}).Delete(&Cover{})

	db.Create(&Cover{
		Url: "https://maximilienproctor.files.wordpress.com/2021/03/house-of-balloons.jpeg", 
		Label: "House of Balloons",
	})

	db.Create(&Cover{
		Url: "https://i.scdn.co/image/ab67616d0000b2738863bc11d2aa12b54f5aeb36",
		Label: "After Hours",
	})

	db.Create(&Cover{
		Url: "https://m.media-amazon.com/images/I/51w6E3vqeyL._AC_SL1200_.jpg",
		Label: "Call Me If You Get Lost",
	})

	db.Create(&Cover{
		Url: "https://images.genius.com/4640c40bd4cec077ba11e54347624ac7.1000x563x1.jpg",
		Label: "Astroworld",
	})

	db.Create(&Cover{
		Url: "https://pbs.twimg.com/media/FoBOa1CWQAEUxol?format=jpg",
		Label: "Her Loss",
	})

	db.Create(&Cover{
		Url: "https://i1.sndcdn.com/artworks-000088468287-jju25j-t500x500.jpg",
		Label: "Days Before Rodeo",
	})


	db.Create(&Cover{
		Url: "https://e.snmc.io/i/1200/s/6e9cf580f2241ebae2107d636160e53f/4931990",
		Label: "Kissland",
	})

	db.Model(&Cover{}).Where("label = ?", "Her Loss").Delete(&Cover{})

	// set database read only
	db.Exec("PRAGMA read_only = ON;")
}

