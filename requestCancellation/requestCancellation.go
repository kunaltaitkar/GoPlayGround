package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)



// problem statement: I want to stop server side processed request when browser is closed.
// Observation : Do not use c as a context from handler func in gin framework (where c is func Test(c *gin.Context))
//    		 Always use c.Request.Context() function for context

// How to test?
//Step 1: Start server
//Step 2: Hit localhost:3030/test in your browser
//step 3: then close tab or browser
func main() {
	g := gin.Default()
	g.GET("/test", Test)

	log.Fatal(g.Run(":3030"))

}

func Test(c *gin.Context) {
	//do not use c as a context
	//use c.Request.Context()
	myFunc(c.Request.Context())
}

func myFunc(ctx context.Context) {
	fmt.Println("Process is running!")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Request Cancelled!")
			return
		default:
			fmt.Println("processing...")
			// return
		}
	}

}
