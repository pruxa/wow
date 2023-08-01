package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pruxa/wow/internal/services"
	"log"
	"net/http"
	"strconv"
)

var (
	js = services.CreateJobService()
	qs = services.CreateQuotesService()
)

func main() {

	// For implementing the api we will use the Echo framework
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.GET("/", challenge)
	e.POST("/", quote)

	// Start server
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func challenge(c echo.Context) error {
	return c.JSON(http.StatusOK, js.GetCurrentJob())
}

func quote(c echo.Context) error {
	numberStr := c.FormValue("number")
	number, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, struct{ Error string }{"Wrong number"})
	}
	hash := c.FormValue("hash")
	valid, err := js.AcceptJob(number, hash)
	if valid {
		return c.JSON(http.StatusOK, qs.GetRandomQuote())
	}
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, struct{ Error string }{"Old job hash"})
	}
	return c.JSON(http.StatusNotAcceptable, struct{ Error string }{"Wrong number"})
}
