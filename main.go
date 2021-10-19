package main

import (
	"GoHttp"
	"net/http"
)

func main() {
	r := GoHttp.New()
	r.GET("/", func(c *GoHttp.Context) {
		c.HTML(http.StatusOK, "<h1>hello GoHttp</h1>")
	})

	r.GET("/hello", func(c *GoHttp.Context) {
		c.String(http.StatusOK, "hello %s you're at %s \n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *GoHttp.Context) {
		c.JSON(http.StatusOK, GoHttp.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.RUN(":9999")
}
