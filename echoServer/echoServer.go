package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

/*

GET http://localhost:3030/ HTTP/1.1
content-type: application/json

{

}

*/

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

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	e.POST("/test", Test)

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)
		for _, l := range locations {
			if err := json.NewEncoder(c.Response()).Encode(l); err != nil {
				return err
			}
			c.Response().Flush()
			time.Sleep(1 * time.Second)
		}
		return nil
	})

	// Start server
	e.Logger.Fatal(e.Start(":3030"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handler
func Test(c echo.Context) error {
	// c.Response
	return c.JSON(200, "success")
}
