package main

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var seq uint64

func main() {
	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	// Middleware
	// Logger disabled as it affects performance VERY badly
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Leave the logger call to keep dependency but it is off
	e.Logger.SetLevel(log.OFF)

	r := e.Group("/api")

	// Routes
	r.GET("/now", func(c echo.Context) error {
		// Create instance of now
		n := &now{
			Timestamp:             time.Now().Unix(),
			TimestampMicroseconds: time.Now().UnixMicro(),
			RequestID:             seq,
		}
		// Increment request counter
		atomic.AddUint64(&seq, 1)
		// Send response
		return c.JSON(http.StatusOK, n)
	})

	r.GET("/nowstream", func(c echo.Context) error {
		// Create instance of now
		n := &now{
			Timestamp:             time.Now().Unix(),
			TimestampMicroseconds: time.Now().UnixMicro(),
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		// Increment request counter
		atomic.AddUint64(&seq, 1)
		return json.NewEncoder(c.Response()).Encode(n)
	})

	r.POST("/jsonpayload", func(c echo.Context) error {
		// Create instance of jsonPayload
		jsonPayload := new(jsonPayload)
		// Try to parse
		if err := c.Bind(jsonPayload); err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		// Increment request counter
		atomic.AddUint64(&seq, 1)
		return c.JSON(http.StatusOK, nil)
	})

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

type jsonPayload struct {
	ID   uint64 `json:"ID"`
	Data string `json:"Data"`
}

type now struct {
	Timestamp             int64  `json:"Timestamp"`
	TimestampMicroseconds int64  `json:"TimestampMicroseconds"`
	RequestID             uint64 `json:"RequestID"`
}
