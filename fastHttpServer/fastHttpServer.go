package main

import (
	"fmt"
	"log"
	"net"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

/*

POST http://localhost:3030/test HTTP/1.1
content-type: application/json

{

}

*/

func main() {
	r := routing.New()

	server := fasthttp.Server{
		Handler: r.HandleRequest,
	}

	var err error

	ln, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started on ", ln.Addr())

	// call for each route
	r.Use(customMethod1, customMethod2)

	//handlers
	r.Post("/test", TestHandler)

	err = server.Serve(ln)
	if err != nil {
		log.Fatal(err)
	}

}

func customMethod1(ctx *routing.Context) error {
	fmt.Println("Custom Method 1 called!")
	return nil
}

func customMethod2(ctx *routing.Context) error {
	fmt.Println("Custom Method 2 called!")
	return nil
}

func TestHandler(c *routing.Context) error {

	fmt.Println("Route called")
	c.SetStatusCode(200)
	return nil
}
