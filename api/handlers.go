package api

import (
	"net/http"
	"strings"

	"github.com/ParasHarnagle/verve_task/models"
	"github.com/ParasHarnagle/verve_task/redis"
	"github.com/gin-gonic/gin"
)

func GetPromotionHandler(c *gin.Context) {
	id := c.Param("id")
	id = strings.ToLower(id)
	p, err := redis.GetPromotionFromCache(id)
	if err != nil {
		p, err = models.GetPromotionFromDatabase(id)
		if err != nil {
			if err.Error() == "promotions.csv is not present, kindly upload it" {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusNotFound, gin.H{"error": "promotion not found"})
			}
			return
		}

		err = redis.PromotionToCache(p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error caching promotion"})
			return
		}
	}
	c.JSON(http.StatusOK, p)
}
