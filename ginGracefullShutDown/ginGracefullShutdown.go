package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*

POST http://localhost:3030/stop HTTP/1.1
content-type: application/json

{
	"name":"kunal"
}

*/

var Quite = make(chan bool)

func main() {

	g := gin.Default()

	g.POST("/stop", stop)

	server := &http.Server{
		Addr:    ":3030",
		Handler: g,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error while starting server: %s\n", err)
		}
	}()

	<-Quite

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("error while shutdown server:", err)
	}
}

func stop(c *gin.Context) {
	Quite <- true
	c.JSON(200, "SUCCESS")
}
