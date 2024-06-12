package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ParasHarnagle/verve_task/api"
	"github.com/ParasHarnagle/verve_task/models"
	"github.com/ParasHarnagle/verve_task/redis"

	//"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

var cleanupMutex sync.Mutex

func main() {
	redis.InitRedis()
	r := gin.Default()
	r.GET("/promotions/:id", api.GetPromotionHandler)

	fmt.Println("Server on localhost:1321")
	go func() {
		log.Fatal(r.RunTLS(":1321", "certs/cert.pem", "certs/key.pem"))
	}()
	cleanupInterval := 30 * time.Minute

	go func() {
		for {
			time.Sleep(cleanupInterval)
			cleanup()
		}
	}()

	select {}
}

// cleanup cleans the immutable file after every 30 min
func cleanup() {
	cleanupMutex.Lock()
	defer cleanupMutex.Unlock()

	redis.ClearCache()

	err := models.DelFile()
	if err != nil {
		log.Println("error deleting csv file: %v", err)
	}
}
