package gincinema

import (
	cinemastorage "ControlMovieTicket/Model/Cinema/Storage"
	cinemabusine "ControlMovieTicket/Model/Cinema/busine"
	cinemamodel "ControlMovieTicket/Model/Cinema/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateCinema(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data cinemamodel.Cinema
		if err := c.ShouldBind(&data); err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := cinemastorage.NewSQLStore(db)
		biz := cinemabusine.NewCinemaBiz(store)
		if err := biz.CreateCinema(c.Request.Context(), &data); err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSONP(http.StatusOK, gin.H{
			"data": data,
		})
	}

}
