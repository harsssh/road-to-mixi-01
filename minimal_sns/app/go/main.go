package main

import (
	"database/sql"
	"net/http"
	"problem1/configs"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := configs.Get()

	db, err := sql.Open(conf.DB.Driver, conf.DB.DataSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "minimal_sns_app")
	})

	e.GET("/get_friend_list", func(c echo.Context) error {
		return c.String(http.StatusOK, "get_friend_list response")
	})

	e.GET("/get_friend_of_friend_list", func(c echo.Context) error {
		return c.String(http.StatusOK, "get_friend_of_friend_list response")
	})

	e.GET("/get_friend_of_friend_list_paging", func(c echo.Context) error {
		return c.String(http.StatusOK, "get_friend_of_friend_list_paging response")
	})

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
}
