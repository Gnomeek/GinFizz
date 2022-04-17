package main

import (
	ginFizz "fizz"
	"net/http"
)

func main() {
	r := ginFizz.New()
	r.GET("/", func(c *ginFizz.Context) {
		c.HTML(http.StatusOK, "<h1>Hello</h1>")
	})

	r.GET("/string", func(c *ginFizz.Context) {
		c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	})

	r.GET("/json", func(c *ginFizz.Context) {
		res := ginFizz.H{}
		for k, v := range c.Request.Header {
			res[k] = v
		}
		c.JSON(http.StatusOK, &res)
	})

	r.POST("/form", func(c *ginFizz.Context) {
		c.JSON(http.StatusOK, &ginFizz.H{
			"username": c.FormValue("username"),
			"password": c.FormValue("password"),
		})
	})

	r.Run(":9999")
}
