package ControlMovieTicket

import (
	cinemamodel "ControlMovieTicket/Model/Cinema/model"
	"ControlMovieTicket/Model/Cinema/transport/gincinema"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	connectString := os.Getenv("CTS")
	db, err := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db, err)
	r := gin.Default()
	v1 := r.Group("/v1")
	Cinema := v1.Group("/cinemas")

	Cinema.GET("/name", func(c *gin.Context) {
		name := c.Param("name")
		var data cinemamodel.Cinema
		db.Where("name=?", name).First(&data)
		c.JSONP(http.StatusOK, gin.H{
			"data": data,
		})
	})

	Cinema.POST("", gincinema.CreateCinema(db))
	Cinema.DELETE("name", func(c *gin.Context) {
		name := c.Param("name")
		var data cinemamodel.Cinema
		db.Delete(&data).Where("name=?", name)
		c.JSONP(http.StatusOK, gin.H{
			"data": "",
		})
	})
}
