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

	r.GET("/hello/:name", func(c *ginFizz.Context) {
		if param, err := c.Param("name"); err != nil {
			c.String(http.StatusBadRequest, "no param %s, err %v", param, err)
		} else {
			c.String(http.StatusOK, "hello %s, you're at %s\n", param, c.Path)
		}

	})

	r.GET("/assets/*filepath", func(c *ginFizz.Context) {
		if param, err := c.Param("filepath"); err != nil {
			c.String(http.StatusBadRequest, "no param %s, err %v", param, err)
		} else {
			c.JSON(http.StatusOK, &ginFizz.H{"filepath": param})
		}
	})

	r.Run(":9999")
}
