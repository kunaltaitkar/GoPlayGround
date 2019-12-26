package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	Geolocation struct {
		Altitude  float64
		Latitude  float64
		Longitude float64
	}
)

var (
	locations = []Geolocation{
		{-97, 37.819929, -122.478255},
		{1899, 39.096849, -120.032351},
		{2619, 37.865101, -119.538329},
		{42, 33.812092, -117.918974},
		{15, 37.77493, -122.419416},
	}
)

type Input struct {
	Name    string `json:"name"`
	LoginID int    `json:"loginId"`
}

func main() {
	g := gin.Default()

	g.POST("/find", Find)

	//stream data
	g.GET("/stream", func(c *gin.Context) {
		chanStream := make(chan int, 10)
		go func() {
			defer close(chanStream)
			for i := 0; i < 5; i++ {
				chanStream <- i
				time.Sleep(time.Second * 1)
			}
		}()
		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-chanStream; ok {
				c.SSEvent("message", msg)
				return true
			}
			return false
		})
	})

	srv := &http.Server{
		Addr:    ":3030",
		Handler: g,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}

/*

POST http://localhost:3030/find HTTP/1.1
content-type: application/json

{
	"name":"kunal"
}

*/

func Find(c *gin.Context) {
	data := Input{}
	c.Bind(&data)

	c.JSON(200, data)

}
