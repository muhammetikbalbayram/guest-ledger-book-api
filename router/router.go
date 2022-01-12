package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"guest-ledger-book-api/collection"
	"log"
	"net/http"
)

func Router() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/guests", collection.GetAllGuests)
	e.POST("/guests", collection.CreateNewGuest)

	if err := e.Start(":5000"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
