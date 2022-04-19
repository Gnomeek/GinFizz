package main

import (
	ginFizz "fizz"
	"log"
	"net/http"
	"time"
)

func onlyForV2() ginFizz.HandlerFunc {
	return func(c *ginFizz.Context) {
		// Start timer
		t := time.Now()
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func main() {
	r := ginFizz.New()
	r.Use(ginFizz.Logger())
	r.GET("/", func(c *ginFizz.Context) {
		c.HTML(http.StatusOK, "<h1>Hello</h1>")
	})
	r.GET("/panic", func(c *ginFizz.Context) {
		names := []string{"ginfizz"}
		c.String(http.StatusOK, names[100])
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
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *ginFizz.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *ginFizz.Context) {
			// expect /hello?name=ginFizzktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.POST("/login", func(c *ginFizz.Context) {
			c.JSON(http.StatusOK, &ginFizz.H{
				"username": c.FormValue("username"),
				"password": c.FormValue("password"),
			})
		})
	}
	r.Run(":9999")
}
