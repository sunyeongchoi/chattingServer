package main

import (
	"chatting_back/pkg/api/login"
	login2 "chatting_back/pkg/login"
	"chatting_back/pkg/websocket"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func serveWS(c echo.Context) error {
	fmt.Println("WebSocket Endpoint Hit")
	pool := websocket.NewPool()
	go pool.Start()
	conn, err := websocket.Upgrade(c.Response(), c.Request())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer conn.Close()

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
	return nil
}

func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/ws", serveWS)
	e.POST("/login", login.GetLoginAPIManager(login2.NewKeycloak()).Login)

	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
